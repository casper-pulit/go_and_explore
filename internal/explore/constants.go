package explore

import (
	"fmt"
)

const s string = "Constant string"

func Constants() {
	fmt.Println(s)

	const pi = 3.14

	const mult int = 2

	fmt.Println(pi, "x", mult, "=", pi*float64(mult))

	tau := pi * float64(mult)

	fmt.Println(tau > pi)

}
