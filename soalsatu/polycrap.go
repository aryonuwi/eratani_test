package soalsatu

func Polycarp(k int) int {
	currentNumber := 1
	count := 0
	for {
		if currentNumber%3 != 0 && currentNumber%10 != 3 {
			count++
		}
		if count == k {
			return currentNumber
		}
		currentNumber++
	}
}
