package main

import (
	"fmt"

	"github.com/aryonuwi/eratani_test/soaldua"
	"github.com/aryonuwi/eratani_test/soalsatu"
)

func main() {
	// soal no 1
	intData := [11]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	for _, val := range intData {
		result := soalsatu.Polycarp(val)
		fmt.Println(result)
	}
	fmt.Println("+============================================+")
	// soal no 2
	testDataPalindrome := [3]interface{}{"kodokkodok", "jinggaaku", 123321}
	for _, val := range testDataPalindrome {
		palindrome := soaldua.Plindrome(val)
		fmt.Println(palindrome)
	}
	fmt.Println("+============================================+")

}
