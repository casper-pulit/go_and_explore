package main

import "fmt"

func main() {
	// declare outside of loop
	i := 1
	// for condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// decalre inside loop; condition; after
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// for range declared in loop
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for without a condition will loop infinetly until broken
	// similar to while true
	for {
		fmt.Println("loop")
		break
	}

	// continue to next iteration of loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
