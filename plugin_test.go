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
	checkResult := monitoringplugin.NewDefaultCheckResult(nil)

	// Do something

	magicNumber := 100
	checkResult.SetPerformanceData("foo", monitoringplugin.NumberUnit(magicNumber))

	if someCheck.critical.Check(float64(magicNumber)) {
		checkResult.SetResult(monitoringplugin.CRTIICAL, "Hello World!")
	} else if someCheck.warning.Check(float64(magicNumber)) {
		checkResult.SetResult(monitoringplugin.WARNING, "Hello World!")
	} else {
		checkResult.SetResult(monitoringplugin.OK, "Hello World!")
	}

	return checkResult
}

func Example() {
	check := new(SomeCheck)
	plugin := monitoringplugin.NewPlugin(check)
	defer plugin.Exit()
	plugin.Start()
	// Output:
	// Hello World! | 'foo'=100;;;0;
}
