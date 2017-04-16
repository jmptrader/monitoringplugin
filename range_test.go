package monitoringplugin

import (
	"fmt"
)

func ExampleRange() {
	r, err := NewRange("@-100:55.00045")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
		return
	}

	fmt.Println(r.ToString())
	fmt.Println(r.Check(50))
	fmt.Println(r.Check(-100.5))
	// Output:
	// @-100:55.00045
	// true
	// false
}

func ExampleRange_infinite() {
	r, err := NewRange("~:0.0")
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

func ExampleRange_zero() {
	r, err := NewRange("@100")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
	}

	fmt.Println(r.ToString())
	fmt.Println(r.Check(200))
	fmt.Println(r.Check(10))
	fmt.Println(r.Check(-1))
	// Output:
	// @100
	// false
	// true
	// false
}

func ExampleRange_empty() {
	r, err := NewRange("")
	if err != nil {
		fmt.Printf("Error can't create range: %s", err)
	}

	fmt.Println(r.ToString())
	fmt.Println(r.Check(200))
	fmt.Println(r.Check(10))
	fmt.Println(r.Check(-1))
	// Output:
	//
	// false
	// false
	// false
}
