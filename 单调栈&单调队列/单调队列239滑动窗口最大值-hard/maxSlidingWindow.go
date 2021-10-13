package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var q [][]int
	var res []int

	lower, higher := 0, k-1
	for j := 0; j < len(nums); j++ {
		//维护一个单调队列->双端队列: 队首和队尾都可以实现出队和入队?
		for len(q) >= 1 && q[len(q)-1][1] <= nums[j] {
			q = q[:len(q)-1]
		}
		q = append(q, []int{j,nums[j]})

		if j == higher {
			for len(q) >= 1 {
				//观察队首元素是否在要求的窗口内
				if q[0][0] <= higher && q[0][0] >= lower {
					res = append(res, q[0][1])
					break
				} else{
					q = q[1:]
				}
			}
			higher++
			lower++
		}
	}
	return res
}

func main(){
	nums := []int{1,-1} //[]int{1,3,-1,-3,5,3,6,7}
	k := 1
	fmt.Println(maxSlidingWindow(nums,k))
}
