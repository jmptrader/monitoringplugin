package monitoringplugin

import (
	"fmt"
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
	FloatPrecision      int
}

type Plugin struct {
	result              CheckResult
	timeout             time.Duration
	timeoutMessage      string
	check               Check
	performanceDataSpec []PerformanceDataSpec
	exited              bool
	floatPrecision      int
}

func NewPlugin(options PluginOpt) (plugin *Plugin) {
	plugin = new(Plugin)
	plugin.exited = false

	if options.Timeout > 0 {
		plugin.timeout = options.Timeout
	} else {
		plugin.timeout = time.Duration(60) * time.Second
	}

	if options.TimeoutMessage != "" {
		plugin.timeoutMessage = options.TimeoutMessage
	} else {
		plugin.timeoutMessage = "Plugin timed out"
	}

	fallbackMessage := options.FallbackMessage
	if fallbackMessage == "" {
		fallbackMessage = "There is no result for this check!"
	}
	plugin.result = NewEasyResult(UNKNOWN, fallbackMessage)

	plugin.check = options.Check

	plugin.performanceDataSpec = options.PerformanceDataSpec

	if options.FloatPrecision < 0 {
		plugin.floatPrecision = 0
	} else if options.FloatPrecision == 0 {
		plugin.floatPrecision = 2
	} else {
		plugin.floatPrecision = options.FloatPrecision
	}

	return
}

func (plugin *Plugin) Start() {
	resultChan := make(chan CheckResult, 1)

	go func() {
		resultChan <- plugin.check.Run()
	}()

	select {
	case result := <-resultChan:
		plugin.result = result
		plugin.Exit()
	case <-time.After(plugin.timeout):
		plugin.result = NewEasyResult(UNKNOWN, plugin.timeoutMessage)
		plugin.Exit()
	}
}

func (plugin *Plugin) floatToStringOrEmpty(value float64) string {
	if value == 0.0 {
		return ""
	}
	return strconv.FormatFloat(value, 'f', plugin.floatPrecision, 64)
}

func (plugin *Plugin) Exit() {
	if plugin.exited {
		return
	}
	plugin.exited = true

	exitCode, message := plugin.result.GetStatus()
	defer os.Exit(exitCode)
	fmt.Print(message)

	perfData := plugin.result.GetPerformanceData()
	if len(plugin.performanceDataSpec) > 0 {
		fmt.Print(" |")
		for _, spec := range plugin.performanceDataSpec {
			var (
				realValue string
			)
			value, hasValue := perfData[spec.Label]
			if hasValue {
				realValue = strconv.FormatFloat(value, 'f', plugin.floatPrecision, 64)
			} else {
				realValue = "U"
			}

			fmt.Printf(" '%s'=%s%s;%s;%s;%s;%s",
				spec.Label, realValue, UnitToString(spec.UnitOfMeasurement),
				spec.Warning.ToString(), spec.Critical.ToString(),
				plugin.floatToStringOrEmpty(spec.Minimum), plugin.floatToStringOrEmpty(spec.Maximum))
		}
		fmt.Print("\n")
	}

	fmt.Print(plugin.result.GetLongOutput())
}
