package main

import (
	"fmt"
	"sort"
)

/*
Given an unsorted integer array nums, return the smallest missing positive integer.
 */

func main() {
	nums := []int{1,2,0}
	nums2 := []int{3,4,-1,1}
	nums3 := []int{7,8,9,11,12}
	nums4 := []int{2,3,1,11,12}
	MissingPosInt(nums)
	MissingPosInt(nums2)
	MissingPosInt(nums3)
	MissingPosInt(nums4)
}

func MissingPosInt (arr []int){
	fmt.Println("Input array to revise: " , arr)
	var result int = 0
	sort.Ints(arr)
	initial := arr[0]
	final := arr[len(arr)-1]
	if initial >1 {
		result = 1
	} else {
		for i := initial; i <= final+1; i++{
			for j := 0; j < len(arr); j++{
				if i == arr[j] {
					break
				} else if j == len(arr)-1 {
					if i > 0{
						result = i
						fmt.Println("The smallest missing positive integer: ", result)
						return
					}
				}
			}
		}
	}
	fmt.Println("The smallest missing positive integer: ", result)
}



