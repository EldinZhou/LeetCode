package main

import "fmt"

func removeDuplicateLetters(s string) string {
	//set是这里面最巧的部分
	//set表示的是某个字母存在于set中
	count, set := make(map[byte]int), make(map[byte]bool)
	var stack []byte

	for _, ch := range []byte(s) {
		count[ch]++
	}

	for _, ch := range []byte(s) {
		count[ch]--
		if !set[ch] {
			for len(stack) >= 1 && stack[len(stack)-1] > ch && count[stack[len(stack)-1]] > 0 {
				set[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1]
			}
			set[ch] = true
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

func main() {

	fmt.Println(removeDuplicateLetters("bbcaac"))
}
