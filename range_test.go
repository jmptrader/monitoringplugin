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
	fmt.Println(r.Check(50))
	fmt.Println(r.Check(-100.5))
	// Output:
	// @-100:55.00045
	// true
	// false
}

func ExampleRangeInfinite() {
	r, err := monitoringplugin.NewRange("~:0.0")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
	}

	fmt.Println(r.ToString())
	fmt.Println(r.Check(-10000))
	fmt.Println(r.Check(10))
	// Output:
	// ~:0
	// false
	// true
}
