package examples

import (
	"fmt"
)

// An interface value that holds a nil concrete value is itself non-nil

func Example_nil_as_interface_argument() {
	var x *some // nil

	fmt.Printf("1. Value: %v, type: %T\n", x, x)

	if x == nil {
		fmt.Println("1. Nil found")
	} else {
		fmt.Println("1. Nil is not found")
	}

	fmt.Println()

	bar(x)

	// Output:
	// 1. Value: <nil>, type: *examples.some
	// 1. Nil found
	//
	// 2. Value: <nil>, type: *examples.some
	// 2. Nil is not found
}

func bar(y explainer) {
	fmt.Printf("2. Value: %v, type: %T\n", y, y)

	if y == nil {
		fmt.Println("2. Nil found")
	} else {
		fmt.Println("2. Nil is not found")
	}
}
