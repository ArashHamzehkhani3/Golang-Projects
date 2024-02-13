package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	var target int
	fmt.Println("Enter a number:")
	fmt.Scan(&target)

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Println(target, index)
	} else {
		fmt.Println("Not found!")
	}
}
