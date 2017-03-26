package monitoringplugin_test

import (
	"fmt"
	"time"

	"github.com/jabdr/monitoringplugin"
)

func ExampleUnit() {
	units := []monitoringplugin.Unit{
		monitoringplugin.DurationUnit(640.0 * time.Second),
		monitoringplugin.DurationUnit((5*86400 + 6*3600 + 10) * time.Second),
		monitoringplugin.CounterUnit(130),
		monitoringplugin.PercentageUnit{Base: 100.0, Quantitiy: 64.0},
		monitoringplugin.BytesUnit(5*1024*1024*1024 + 10*1024*1024),
	}

	for index, unit := range units {
		fmt.Printf("Test %d:\n", index)
		fmt.Printf("%.2f%s\n", unit.Value(), unit.UnitString())
		fmt.Println(unit.String())
		fmt.Println("")
	}

	// Output:
	// Test 0:
	// 640.00s
	// 10m40s
	//
	// Test 1:
	// 453610.00s
	// 126h0m10s
	//
	// Test 2:
	// 130.00c
	// 130c
	//
	// Test 3:
	// 64.00%
	// 64.00%
	//
	// Test 4:
	// 5379194880.00B
	// 5.01GiB
}
