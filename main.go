package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{16, 4, 2, 6, 10, 8, 1}))
	fmt.Println(bubbleRecursion([]int{16, 4, 2, 6, 10, 8, 1}))
}

func bubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		last := len(list) - i
		for j := 1; j < last; j++ {
			if list[j-1] > list[j] {
				swap(list, j-1, j)
			}
		}
	}
	return list
}

func bubbleRecursion(list []int) []int {
	round := 0
	bubbleRecursionHelper1(list, round)
	return list
}

func bubbleRecursionHelper1(list []int, round int) {
	if round >= len(list) {
		return
	}
	last := len(list) - round
	bubbleRecursionHelper2(list, last, 1)
	bubbleRecursionHelper1(list, round+1)
}

func bubbleRecursionHelper2(list []int, len, iRun int) {
	if iRun >= len {
		return
	}
	if list[iRun-1] > list[iRun] {
		swap(list, iRun-1, iRun)
	}
	bubbleRecursionHelper2(list, len, iRun+1)
}

func swap(nums []int, iLeft, iRight int) {
	temp := nums[iLeft]
	nums[iLeft] = nums[iRight]
	nums[iRight] = temp
}
