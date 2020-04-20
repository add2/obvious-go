package examples

import (
	"fmt"
)

// The argument passed to the function as an interface can be either a value or a pointer

func Example_passing_interface_by_pointer() {
	var x some
	var y *some

	foo(x, y)

	// Output:
	// x = 0
	// y = <nil>
}

func foo(x explainer, y explainer) {
	fmt.Println("x =", x) // This is value. Can't be nil.
	fmt.Println("y =", y) // This is pointer. Can be nil. It is not obvious if we consider only the code of this function.
}
