package backtrack

import (
	"fmt"
)

var matched = false

func regular(pattern string, text string) bool {
	match(0, 0, pattern, text)
	return matched
}

func match(pi int, tj int, pattern string, text string){
	if matched {
		return
	}
	if pi == len(pattern) {
		if tj == len(text) {
			fmt.Println()
			matched = true
		}
		return
	}
	if pattern[pi] == '*' {
		for k:=0; k<len(text)-tj; k++ {
			match(pi+1, tj+k, pattern,text)
		}
	} else if pattern[pi] == '?' {
		match(pi+1, tj, pattern, text)
		match(pi+1, tj+1, pattern,text)
	} else if tj < len(text) && pattern[pi] == text[tj] {
		match(pi+1, tj+1, pattern,text)
	}
}

/*
func main()  {
	pattern := "a*b?c"
	text := "accccccccccbc"
	ret := regular(pattern, text)
	fmt.Println(ret)
}
 */