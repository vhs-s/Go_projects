package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	min_value := min(a, b)
	max_value := max(a, b)
	for max_value != min_value {
		if max_value-min_value > min_value {
			max_value = max_value - min_value
		} else {
			temp := min_value
			min_value = max_value - min_value
			max_value = temp
		}
	}
	fmt.Println(min_value)
}
