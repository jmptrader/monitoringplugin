package monitoringplugin_test

import (
	"flag"
	"math"
	"time"

	"github.com/jabdr/monitoringplugin"
)

type SomeCheck struct {
	warning  monitoringplugin.Range
	critical monitoringplugin.Range
}

func (someCheck SomeCheck) HandleArguments(options monitoringplugin.PluginOpt) (monitoringplugin.PluginOpt, error) {
	warning := flag.String("w", "", "Warning range")
	critical := flag.String("c", "", "Critical range")
	timeout := flag.Int64("t", 60, "Plugin timeout in seconds")
	flag.Parse()

	warnRange, err := monitoringplugin.NewRange(*warning)
	if err != nil {
		return options, err
	}
	someCheck.warning = warnRange

	critRange, err := monitoringplugin.NewRange(*critical)
	if err != nil {
		return options, err
	}
	someCheck.critical = critRange

	options.Timeout = time.Duration(*timeout) * time.Second
	options.Check = someCheck

	options.PerformanceDataSpec = []monitoringplugin.PerformanceDataSpec{
		{
			Label:             "foo",
			UnitOfMeasurement: monitoringplugin.NumberUnitSpecification,
			Warning:           &warnRange,
			Critical:          &critRange,
			Minimum:           0,
			Maximum:           math.Inf(1),
		},
	}

	options.DoNotExit = true // for testing

	return options, nil
}

func (someCheck SomeCheck) Run() monitoringplugin.CheckResult {
	// Do some foo

	return monitoringplugin.NewEasyResult(monitoringplugin.OK, "Everything is fine")
}

func Example() {
	check := new(SomeCheck)
	plugin := monitoringplugin.NewPlugin(check)
	defer plugin.Exit()
	plugin.Start()
	// Output:
	// Everything is fine | 'foo'=U;;;0;
}
