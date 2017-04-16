package monitoringplugin

import (
	"fmt"
	"time"
)

func ExampleUnit() {
	units := []Unit{
		DurationUnit(640.0 * time.Second),
		DurationUnit((5*86400 + 6*3600 + 10) * time.Second),
		CounterUnit(130),
		PercentageUnit{Base: 100.0, Quantitiy: 64.0},
		BytesUnit(5*1024*1024*1024 + 10*1024*1024),
	}

	for index, unit := range units {
		fmt.Printf("Test %d:\n", index)
		fmt.Printf("%.2f\n", unit.Value())
		fmt.Println(unit.HumanReadable())
		fmt.Println("")
	}

	// Output:
	// Test 0:
	// 640.00
	// 10m40s
	//
	// Test 1:
	// 453610.00
	// 126h0m10s
	//
	// Test 2:
	// 130.00
	// 130c
	//
	// Test 3:
	// 64.00
	// 64.00%
	//
	// Test 4:
	// 5379194880.00
	// 5.01GiB
}
