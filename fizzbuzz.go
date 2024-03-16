package main

import "fmt"

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

func main() {

	for i := 1; i <= 10; i++ {
		odd_or_even(i)
	}

	for i := 1; i <= 100000; i++ {
		fizzBuzz(i)
	}

}
