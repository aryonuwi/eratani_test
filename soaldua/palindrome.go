package soaldua

import (
	"fmt"
	"strings"
)

func Plindrome(data interface{}) string {
	strConvert := fmt.Sprintf("%v", data)
	str := strings.ToLower(strings.ReplaceAll(strConvert, " ", ""))
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return "bukan palindrome"
		}
	}
	return "itu palindrome"
}
