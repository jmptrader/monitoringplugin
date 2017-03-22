package monitoringplugin_test

import (
	"fmt"
	"math"

	"github.com/jabdr/monitoringplugin"
)

func ExamplePerformanceDataSpec() {
	warnRange, err := monitoringplugin.NewRange("80")
	if err != nil {
		fmt.Println(err)
	}
	critRange, err := monitoringplugin.NewRange("90")
	if err != nil {
		fmt.Println(err)
	}
	specs := []monitoringplugin.PerformanceDataSpec{
		{
			Label:             "example",
			UnitOfMeasurement: monitoringplugin.SECONDS,
			Minimum:           0.0,
			Maximum:           math.Inf(1),
			Warning:           &warnRange,
			Critical:          &critRange,
		},
		{
			Label:             "counter",
			UnitOfMeasurement: monitoringplugin.COUNTER,
			Maximum:           math.Inf(1),
		},
		{
			Label:             "novalue",
			UnitOfMeasurement: monitoringplugin.MEGABYTES,
			Minimum:           math.Inf(-1),
			Maximum:           math.Inf(1),
		},
	}

	perfData := map[string]float64{
		"example": 40.5,
		"counter": 60,
	}

	for _, spec := range specs {
		fmt.Println(spec.FormatPerfDataFromMap(perfData))
	}

	// Output:
	// 'example'=40.5s;80;90;0;
	// 'counter'=60c;;;0;
	// 'novalue'=UMB;;;;
}
