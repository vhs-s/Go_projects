package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var m int
	fmt.Scan(&m)
	for i := 0; i < 20; i++ {
		tim := time.Now()
		selectionSort(randomizeArray(m))
		fmt.Println("-----------------SORTED---------------------", "i = ", i)
		elapsed := time.Since(tim)
		fmt.Println(elapsed)
	}

}

func selectionSort(nums []float64) []float64 {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				count++
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	fmt.Println(count)
	return nums

}

func randomizeArray(m int) []float64 {
	nums := make([]float64, 0, m)
	for i := 0; i < m; i++ {
		nums = append(nums, ((rand.Float64() * 2) - 1))
	}
	return nums
}
