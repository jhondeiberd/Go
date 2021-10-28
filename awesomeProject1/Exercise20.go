package main

import (
	"fmt"
)

/* Exercise2
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. 
All houses at this place area r ranged in a circle.That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two
adjacent houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house, return the maximum amount
of money you can rob tonight without alerting the police.
Example 1:
Input: nums = [2,3,2] Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.

Example 2:
Input: nums = [1,2,3,1] Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3). Total amount you can rob = 1 + 3 = 4.

Example 3:
Input: nums = [1,2,3] Output: 3

Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 1000
 */

func main() {
	var houses []int = []int{1,2,3,1}
	var houses1 []int = []int{1, 3, 4, 4, 3, 3, 7, 2, 3, 4, 5, 1}
	var houses2 []int = []int{2,3,2}
	var houses3 []int = []int{1,2,3}
	robber(houses)
	robber(houses1)
	robber(houses2)
	robber(houses3)
}

func robber (arr[] int) {
	var qtyHouses int = len(arr)
	var message string = "The amount of money can to rob tonight is: "
	var message1 string = "The houses to rob: "
	output:=0
	output2:=0
	if qtyHouses == 0 {
		fmt.Println(message1, arr)
		fmt.Println(message, 0)
	} else if qtyHouses == 1 {
		fmt.Println(message1, arr)
		fmt.Println(message, arr[0])
	} else if qtyHouses < 4 {
		height:=0
		for i:=0; i < qtyHouses; i++{
			if arr[i] >= height {
				height = arr[i]
			}
		}
		fmt.Println(message1, arr)
		fmt.Println(message, height)
	} else {
		for i:=0;i< len(arr);i=i+2{
			output=arr[i]+output
		}
		for i:=1;i<len(arr);i=i+2{
			output2=arr[i]+output2
		}
		if output >output2{
			fmt.Println(message1, arr)
			fmt.Println(message, output)
		}else {
			fmt.Println(message1, arr)
			fmt.Println(message, output)
		}
	}
}



