package main

import (
	"fmt"

	"github.com/aryonuwi/eratani_test/soalSatu"
)

func main() {
	// soal no 1
	k := [11]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	for _, val := range k {
		result := soalSatu.Polycarp(val)
		fmt.Println(result)
	}

}
