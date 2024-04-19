package main

import (
	"fmt"
	// imported modules external or internal must be specified in go.mod
	"github.com/casper-pulit/go_and_explore/internal/explore"
)

func main() {
	fmt.Println("Main!")
	// exported functions must be Capitalised
	fmt.Println(explore.HelloWorld())
	explore.Values()
	explore.Variables()
}
