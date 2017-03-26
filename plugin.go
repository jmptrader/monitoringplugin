package monitoringplugin

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type PluginOpt struct {
	Timeout             time.Duration
	TimeoutMessage      string
	FallbackMessage     string
	Check               Check
	PerformanceDataSpec []PerformanceDataSpec
	DoNotExit           bool
}

type Plugin struct {
	cliHandler          CliHandler
	result              CheckResult
	timeout             time.Duration
	timeoutMessage      string
	check               Check
	performanceDataSpec []PerformanceDataSpec
	exited              bool
	doNotExit           bool
}

type CliHandler interface {
	HandleArguments(options PluginOpt) (PluginOpt, error)
}

func NewPlugin(cli CliHandler) (plugin *Plugin) {
	plugin = new(Plugin)
	plugin.exited = false
	plugin.cliHandler = cli
	return
}

func (plugin *Plugin) handleCli() {
	defaultOptions := PluginOpt{
		Timeout:         time.Duration(60) * time.Second,
		TimeoutMessage:  "Plugin timed out",
		FallbackMessage: "There is no result for this check!",
	}
	options, err := plugin.cliHandler.HandleArguments(defaultOptions)
	if err != nil {
		plugin.result = &EasyResult{UNKNOWN, fmt.Sprintf("Could not handle CLI: '%s'", err)}
		plugin.Exit()
	}

	plugin.timeout = options.Timeout
	plugin.timeoutMessage = options.TimeoutMessage
	plugin.result = &EasyResult{UNKNOWN, options.FallbackMessage}
	plugin.check = options.Check
	plugin.performanceDataSpec = options.PerformanceDataSpec
	plugin.doNotExit = options.DoNotExit
}

func (plugin *Plugin) runCheck() {
	resultChan := make(chan CheckResult, 1)

	go func() {
		resultChan <- plugin.check.Run()
	}()

	select {
	case result := <-resultChan:
		plugin.result = result
		plugin.Exit()
	case <-time.After(plugin.timeout):
		plugin.result = &EasyResult{UNKNOWN, plugin.timeoutMessage}
		plugin.Exit()
	}
}

func (plugin *Plugin) Start() {
	plugin.handleCli()
	plugin.runCheck()
}

func (plugin *Plugin) floatToStringOrEmpty(value float64) string {
	if math.IsInf(value, 1) || math.IsInf(value, -1) {
		return ""
	}
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// Exit is the function that outputs the result of this plugin.
// You should call it with defer.
func (plugin *Plugin) Exit() {
	if plugin.exited {
		return
	}
	plugin.exited = true

	exitCode, message := plugin.result.GetStatus()
	if !plugin.doNotExit {
		defer os.Exit(exitCode)
	}

	fmt.Print(message)

	perfData := plugin.result.GetPerformanceData()
	if len(plugin.performanceDataSpec) > 0 {
		fmt.Print(" |")
		for _, spec := range plugin.performanceDataSpec {
			fmt.Printf(" %s", spec.FormatPerfDataFromMap(perfData))
		}
		fmt.Print("\n")
	}

	fmt.Print(plugin.result.GetLongOutput())
}
