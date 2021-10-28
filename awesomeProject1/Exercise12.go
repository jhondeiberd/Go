package main

import (
	"fmt"
	"strconv"
)

/* Exercise1-2
Given an m x n matrix, return all elements of the matrix in spiral order.
*/

func main() {
	fmt.Println("**** Exercise 1  *****")
	var matrix [][]int
	matrix = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rows := len(matrix)
	columns := len(matrix[0])
	fmt.Print("Matrix origin: \t", matrix, "\n")
	fmt.Print("Spiral print : \t")
	SpiralMatrix(rows,columns,matrix)

	fmt.Println("\n\n**** Exercise 2  *****")
	var matrix2 [][]int
	matrix2 = [][]int{{1,2,3,4},{5,6,7,8,},{9,10,11,12}}
	rows2 := len(matrix2)
	columns2 := len(matrix2[0])
	fmt.Print("Matrix origin: \t", matrix2, "\n")
	fmt.Print("Spiral print : \t")
	SpiralMatrix(rows2,columns2,matrix2)
}

func SpiralMatrix(m int, n int, a[][] int){
	var i, j, k int
	j = 0
	k = 0

	for j < m && k < n {
		for i = k; i < n; i++{
			fmt.Print(strconv.Itoa(a[j][i]) + " ")
		}
		j++

		for i = j; i < m; i++{
			fmt.Print(strconv.Itoa(a[i][n-1]) + " ")
		}
		n--

		if j < m {
			for i = n-1; i >=k; i--{
				fmt.Print(strconv.Itoa(a[m-1][i]) + " ")
			}
			m--
		}

		if k < n {
			for i = m-1; i >= j; i--{
				fmt.Print(strconv.Itoa(a[i][k]) + " ")
			}
			k++
		}
	}
}