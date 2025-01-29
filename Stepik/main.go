package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	converted_int := strconv.Itoa(x)
	var new_str string
	for i := len(converted_int) - 1; i != -1; i-- {
		new_str += string(converted_int[i])
	}
	new_int, _ := strconv.Atoi(new_str)

	if new_int == x {
		return true
	} else {
		return false
	}
}
