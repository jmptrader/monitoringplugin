package monitoringplugin

import (
	"fmt"
	"math"
	"time"
)

func ExamplePerformanceDataSpec() {
	warnRange, err := NewRange("80")
	if err != nil {
		fmt.Println(err)
	}
	critRange, err := NewRange("90")
	if err != nil {
		fmt.Println(err)
	}
	specs := []PerformanceDataSpec{
		{
			Label:             "example",
			UnitOfMeasurement: DurationUnitSpecification,
			Minimum:           0.0,
			Maximum:           math.Inf(1),
			Warning:           &warnRange,
			Critical:          &critRange,
		},
		{
			Label:             "counter",
			UnitOfMeasurement: CounterUnitSpecification,
			Maximum:           math.Inf(1),
		},
		{
			Label:             "novalue",
			UnitOfMeasurement: BytesUnitSpecification,
			Minimum:           math.Inf(-1),
			Maximum:           math.Inf(1),
		},
		{
			Label:             "percenttest",
			UnitOfMeasurement: PercentageUnitSpecification,
			Minimum:           0.0,
			Maximum:           100.0,
		},
	}

	perfData := map[string]Unit{
		"example": DurationUnit(40 * time.Second),
		"counter": CounterUnit(60),
		"percenttest": PercentageUnit{
			Base:      100.0,
			Quantitiy: 45.0,
		},
	}

	for _, spec := range specs {
		fmt.Println(spec.FormatPerfDataFromMap(perfData))
	}

	// Output:
	// 'example'=40s;80;90;0;
	// 'counter'=60c;;;0;
	// 'novalue'=UB;;;;
	// 'percenttest'=45%;;;0;100
}
