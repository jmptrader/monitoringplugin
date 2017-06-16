package monitoringplugin

import "math"

type DynPerfCheck struct {
}

func (someCheck DynPerfCheck) HandleArguments(options PluginOpt) (PluginOpt, error) {
	options.Check = someCheck

	options.PerformanceDataSpec = []PerformanceDataSpec{}

	return options, nil
}

func (someCheck DynPerfCheck) Run() CheckResult {
	checkResult := NewDefaultCheckResult(nil)

	// Do something
	checkResult.SetPerformanceDataSpec(PerformanceDataSpec{
		Label:             "foo",
		Minimum:           0.0,
		Maximum:           math.Inf(1),
		UnitOfMeasurement: NumberUnitSpecification,
	})

	magicNumber := 100
	checkResult.SetPerformanceData("foo", NumberUnit(magicNumber))

	checkResult.SetResult(OK, "Hello World!")

	return checkResult
}

func ExampleDynamicPerfData() {
	check := new(DynPerfCheck)
	plugin := NewPlugin(check)
	defer plugin.Exit()
	plugin.Start()
	// Output:
	// Hello World! | 'foo'=100;;;0;
}
