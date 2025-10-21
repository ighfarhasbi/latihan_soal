package main

import (
	"fmt"
)

func main() {
	res := findSumMin([]int{5, 1, 4}, 4)
	fmt.Println(res)

}

func findSumMin(arr []int, m int) int {
	sum := 0
	for i := 0; i < m; i++ {
		minIdx := findMinIdx(arr)
		if minIdx == -1 {
			break
		}
		sum += arr[minIdx]
		arr[minIdx]--
	}
	return sum
}

func findMinIdx(arr []int) int {
	minIdx := -1
	for i, v := range arr {
		if v <= 0 {
			continue
		}
		if minIdx == -1 || arr[i] < arr[minIdx] {
			minIdx = i
		}
	}
	return minIdx
}
