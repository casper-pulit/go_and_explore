package explore

import (
	"fmt"
)

func Variables() {
	// go will infer type
	var a = "a string"
	fmt.Println(a)

	// multi var declation
	var b, c int = 1, 2
	fmt.Println(b, c)

	// vars declared without initialising value will be zero-d
	var no_init_int int
	fmt.Println("no init int", no_init_int)

	var no_init_bool bool
	fmt.Println("no init bool", no_init_bool)

	// shorthand walrus operator

	walrus := "cute"
	fmt.Println("Warlus := are", walrus)

}
