package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	input: root = [3,9,20,null,null,15,7]
	fmt.Printf("\n\n\n")
	levelOrder(input)

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q = [](*TreeNode){root}
	var ans = [][]int{}
	for len(q) > 0 {
		var curLevel = []int{}
		var sz = len(q)
		for i:=0; i < sz; i++ {
			var cur = q[i]
			curLevel = append(curLevel, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		q = q[sz:]
		ans = append(ans, curLevel)
	}
	return ans
}