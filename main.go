package main

import "fmt"

func main() {
	// Comments are added using double forward-slashes

	// Print == Println
	// fmt library
	fmt.Println("Hello World!")

	fmt.Println("Strings", "can be", "separated", "by ,'s")

	// Str and numeric can be printed together when separated by comma
	fmt.Println("1 + 1 =", 1+1)
	// X.X where X is a number signifies a float
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// 13/3 != 13.0/3.0
	fmt.Println("13/3 =", 13/3)
	fmt.Println("13.0/3.0 =", 13.0/3)
	fmt.Println("13/3 == 13.0/3.0:", 13/3 == 13.0/3.0)
	// No difference whether result of equation of type int calculated with float or int
	fmt.Println("6/2 == 6.0/2.0:", 6/2 == 6.0/2.0)

	// Int division implcitly rounds down (floor)
	fmt.Println("2/3 = 0.66666 but actually:", 2/3)
	fmt.Println("1/3 = 0.33333 but actually:", 1/3)

	// Escaped characters
	fmt.Println("Line 1\nLine 2\nLine3: Tab 1\tTab2\nLine: \\n, \\t\n")

	// Boolean

	fmt.Println("true:", true)
	fmt.Println("false:", false)
	fmt.Println("true and true:", true && true)
	fmt.Println("true and false:", true && false)
	fmt.Println("false and true:", false && true)
	fmt.Println("true or false:", true || false)
	fmt.Println("false or false:", false || false)
	fmt.Println("not true:", !true)
	fmt.Println("not false:", !false)

	// bools don't support mathematical operations

}
