package monitoringplugin

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

var (
	testingMode = false
)

// PluginOpt sets the default values for the Plugin.
// Is is used by the CliHandler implementation.
type PluginOpt struct {
	// Set the timeout for the plugin.
	Timeout time.Duration
	// Set the message for plugin when the timeout has been hit.
	TimeoutMessage string
	// If the check crashes this message will be shown.
	FallbackMessage string
	// The actual check that should be run.
	Check Check
	// The specification for the performance data output (not the values)
	PerformanceDataSpec []PerformanceDataSpec
	// Do not exit the plugin, when plugin.Exit will be called.
	// This is for testing purposes.
	DoNotExit bool
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

// CliHandler is an interface for a type that should parse cli parameter
// and prepare the plugin.
type CliHandler interface {
	HandleArguments(options PluginOpt) (PluginOpt, error)
}

// NewPlugin is a constructor for Plugin
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

// Start will start the CliHandler and run the check.
// You should call defer plugin.Exit() before!
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
	if !plugin.doNotExit && !testingMode {
		defer os.Exit(exitCode)
	}

	fmt.Print(message)

	perfData := plugin.result.GetPerformanceData()
	if len(plugin.performanceDataSpec) > 0 {
		fmt.Print(" |")
		for _, spec := range plugin.performanceDataSpec {
			fmt.Printf(" %s", spec.FormatPerfDataFromMap(perfData))
		}
	}
	fmt.Print("\n")

	fmt.Print(plugin.result.GetLongOutput())
}
