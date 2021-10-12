package main

import "fmt"

func main(){
	//KMP算法的复杂度是 O(m+n) n = len(s), m = len(p)
	//https://www.zhihu.com/question/21923021 Reference
	s, p := "ababaabaabacabaabacabaabac", "abaabac" //s:主串, p:模式串
	//kmp第一步是先获得p模式串的next数组, 即利用next数组来存储p字符串中
	//[0...i]子串中的最大前后缀数,也就是当前子串最大的前缀==后缀(前缀后缀相同)的长度k
	//k != i+1, 不能包含自身, 也就是前缀必须包含首位但不能包含末尾, 反之后缀必须包含末尾不能包含前缀
	//next 数组各值的含义：代表当前字符之前的字符串中，有多大长度的相同前缀后缀。
	//例如，如果 next [j] = k，代表 j 之前的字符串中有最大长度为 k 的相同前缀后缀
	//一、获取next数组
	//fmt.Println(s, p)
	//fmt.Println(getNext("abcabdddabcabc"))
	sidx, pidx, count := 0, 0, 0
	nxt := getNext(p)
	for sidx < len(s) {
		if s[sidx] == p[pidx] {
			sidx++
			pidx++
		} else if pidx != 0 {
			pidx = nxt[pidx-1]
		} else { //pidx退无可退，则舍弃当前的情况继续往下做
			sidx++
		}
		if pidx == len(p) {
			fmt.Printf("Matched Pos: %d\n", pidx)
			count++
			pidx = nxt[pidx-1]
		}
	}
	fmt.Printf("Matched Counts: %d\n", count)
}

//O(m^2)
func getNextBF(x int, p string) int {
	for i := x; i >= 0; i-- {
		if p[0:i] == p[x-i+1:x+1] {
			return i
		}
	}
	return 0
}

//O(m):优化后的next数组, 思想类似动态规划
func getNext(p string) []int {
	next := make([]int, len(p), len(p))
	j, now := 1, 0
	for j < len(p) {
		if p[now] == p[j] {
			now++
			next[j] = now
			j++
		} else if now != 0 {
			now = next[now-1]
		} else {
			j++
		}
	}
	return next
}