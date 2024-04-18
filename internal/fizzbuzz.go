package main

import "fmt"

// Define function (no return)
func odd_or_even(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}

func fizzBuzz(n int) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Println(n, "FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println(n, "Fizz")
	} else if n%5 == 0 {
		fmt.Println(n, "Buzz")
	} else {
		fmt.Println(n)
	}
}

// func with return
// func func_name(arg arg_type, ...) return_type {...}
func sum(a int, b int, c int) int {
	return (a + b + c)
}

// func multiple return values
// func func_name(arg arg_type, ...) (return_type_1, return_type_2) {...}
func vals() (int, int) {
	return 3, 7
}

// variadic functions
// can take any number of arugments
// e.g. fmt.Println

// func func_name(arg ... arg_type, ...) return_type {...}
func betterSum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	for i := 1; i <= 10; i++ {
		odd_or_even(i)
	}

	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}

	var out = sum(1, 2, 3)

	fmt.Println(out)

	fmt.Println("\n\n\n\n")

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// blank identifer "_" to get subset of values
	_, c := vals()
	fmt.Println(c)

	fmt.Println("\n\n\n\n")

	betterSum(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	betterSum(nums...)
}
