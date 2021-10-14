package main

import (
	"fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	//求解思路是我们需要固定高度和宽中的一个，我们选择固定高度，那么就是找到
	//这个高度下最大的宽度，也就可以确定面积
	ans := 0
	n := len(heights)
	var stack, left []int
	for i := 0; i < len(heights); i++ {
		//维护一个单调栈
		//heights := []int{2,1,5,6,2,3}
		for len(stack) >= 1 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) >= 1 {
			left = append(left, stack[len(stack)-1])
		} else {
			left = append(left, -1)
		}
		stack = append(stack, i)
	}
	var stackR []int
	right :=  make([]int, n)
	for j := n-1; j >= 0; j-- {
		for len(stackR) >= 1 && heights[stackR[len(stackR)-1]] >= heights[j] {
			stackR = stackR[:len(stackR)-1]
		}
		//heights := []int{2,1,5,6,2,3}
		if len(stackR) >= 1 {
			right[j] = stackR[len(stackR)-1]
		} else {
			right[j] = n
		}
		stackR = append(stackR, j)
	}

	for k := 0; k < n; k++ {
		ans = int(math.Max(float64(ans), float64((right[k]-left[k]-1)*heights[k])))
	}
	fmt.Println(left, right)
	return ans
}

func main() {
	heights := []int{0,9}
	fmt.Println(largestRectangleArea(heights))
}
