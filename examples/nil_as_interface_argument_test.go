package examples

import (
	"fmt"
)

// A nil pointer passed to a function as an interface does not pass nil verification

func Example_nil_as_interface_argument() {
	var x *some // nil

	if x == nil {
		fmt.Println("1. Nil found")
	} else {
		fmt.Println("1. Nil is not found")
	}

	bar(x)

	// Output:
	// 1. Nil found
	// 2. Nil is not found
}

func bar(x explainer) {
	if x == nil {
		fmt.Println("2. Nil found")
	} else {
		fmt.Println("2. Nil is not found")
	}
}
