package main

import (
	"fmt"

	"github.com/aryonuwi/eratani_test/soaldua"
	"github.com/aryonuwi/eratani_test/soalempat"
	"github.com/aryonuwi/eratani_test/soalsatu"
)

func main() {
	// Soal no 1
	intData := [11]int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	for _, val := range intData {
		result := soalsatu.Polycarp(val)
		fmt.Println(result)
	}

	fmt.Println("+============================================+")
	// Soal no 2
	testDataPalindrome := [3]interface{}{"kodokkodok", "jinggaaku", 123321}
	for _, val := range testDataPalindrome {
		palindrome := soaldua.Plindrome(val)
		fmt.Println(palindrome)
	}

	fmt.Println("+============================================+")
	// Soal no 4
	randomNumbers := []int{9, 2, 5, 1, 7, 3, 8, 4, 6}
	response := soalempat.Sorting(randomNumbers)
	fmt.Println(response)

	fmt.Println("+============================================+")
	// Soal no 3
}
