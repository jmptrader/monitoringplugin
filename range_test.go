package monitoringplugin_test

import (
	"fmt"

	"github.com/jabdr/monitoringplugin"
)

func ExampleRange() {
	r, err := monitoringplugin.NewRange("@-100:55.00045")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
	}

	fmt.Println(r.ToString())
	// Output:
	// @-100:55.00045
}

func ExampleRangeInfinite() {
	r, err := monitoringplugin.NewRange("~:0.0")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
	}

	fmt.Println(r.ToString())
	// Output:
	// ~:0
}
