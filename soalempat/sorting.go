package soalempat

func Sorting(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, middle, right []int
	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			middle = append(middle, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(Sorting(left), middle...), Sorting(right)...)
}
