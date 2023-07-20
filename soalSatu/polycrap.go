package soalSatu

func Polycarp(k int) int {
	// Initialize variables
	currentNumber := 1
	count := 0

	for {
		// Check if the current number meets the conditions
		if currentNumber%3 != 0 && currentNumber%10 != 3 {
			count++
		}

		// If we reach the k-th element, return it
		if count == k {
			return currentNumber
		}

		// Move to the next number
		currentNumber++
	}
}
