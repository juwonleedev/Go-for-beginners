package main

import (
	"fmt"
)

func canIDrink1 (age int) bool {
	if age < 18 {
		return false
	}
	return true
}

// variable expression - if,else에서 사용하기 위해 선언
func canIDrink2 (age int) bool {
	if koreanAge := age + 2 ; koreanAge < 18 {
		return false
	}
	return true 
}

func main () {
	fmt.Println(canIDrink2(16))
}