package monitoringplugin

import (
	"flag"
	"math"
	"time"
)

func init() {
	testingMode = true
}

type SomeCheck struct {
	warning  Range
	critical Range
}

func (someCheck SomeCheck) HandleArguments(options PluginOpt) (PluginOpt, error) {
	warning := flag.String("w", "", "Warning range")
	critical := flag.String("c", "", "Critical range")
	timeout := flag.Int64("t", 60, "Plugin timeout in seconds")
	flag.Parse()

	warnRange, err := NewRange(*warning)
	if err != nil {
		return options, err
	}
	someCheck.warning = warnRange

	critRange, err := NewRange(*critical)
	if err != nil {
		return options, err
	}
	someCheck.critical = critRange

	options.Timeout = time.Duration(*timeout) * time.Second
	options.Check = someCheck

	options.PerformanceDataSpec = []PerformanceDataSpec{
		{
			Label:             "foo",
			UnitOfMeasurement: NumberUnitSpecification,
			Warning:           &warnRange,
			Critical:          &critRange,
			Minimum:           0,
			Maximum:           math.Inf(1),
		},
	}

	return options, nil
}

func (someCheck SomeCheck) Run() CheckResult {
	checkResult := NewDefaultCheckResult(nil)

	// Do something

	magicNumber := 100
	checkResult.SetPerformanceData("foo", NumberUnit(magicNumber))

	if someCheck.critical.Check(float64(magicNumber)) {
		checkResult.SetResult(CRITICAL, "Hello World!")
	} else if someCheck.warning.Check(float64(magicNumber)) {
		checkResult.SetResult(WARNING, "Hello World!")
	} else {
		checkResult.SetResult(OK, "Hello World!")
	}

	return checkResult
}

func Example() {
	check := new(SomeCheck)
	plugin := NewPlugin(check)
	defer plugin.Exit()
	plugin.Start()
	// Output:
	// Hello World! | 'foo'=100;;;0;
}
