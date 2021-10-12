package main

import "fmt"

func main(){
	s, p := "AAAAAABC", "AAAB" //s:主串, p:模式串
	count := 0
	//暴力解法的复杂度是O(mn) n = len(s), m = len(p)
	for i := 0; i <= len(s) - len(p) ; i++ {
		if s[i:i+len(p)] == p { //这个也需要逐位比较
			fmt.Printf("Matched Pos: %d\n", i) //匹配位置
			count++
		}
	}
	fmt.Printf("Matched Counts: %d\n", count) //匹配次数
}
