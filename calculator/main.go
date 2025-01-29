package main

import (
	"fmt"
)

func main() {
	var allowedsymbols = []rune("+-*/")
	var a, b int
	var usersymbol string
	var result any
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&usersymbol)
	if checktoallowedsymbol(allowedsymbols, usersymbol) {
		result = calculate(a, b, usersymbol)
		if result != nil {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Неверная операция")
	}
}

func checktoallowedsymbol(symbols []rune, usersym string) bool {
	for i := 0; i < len(symbols); i++ {
		if string(symbols[i]) == usersym {
			return true
		}
	}
	return false
}

func calculate(a, b int, usersym string) any {
	switch usersym {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		var result float64
		result = float64(a) / float64(b)
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return nil
		}
		return result
	}

	return 0
}
