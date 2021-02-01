package kmp

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	str := "bbc abcdab abcdaaaacaajaaaajaajbcdabde";
	mark := "aaaajaaj"
	//
	//mark = "abcabcacab"
	//str = "bbc abcabcabcacabdab abcdabcabcacab"
	//
	//mark = "abc"
	//str = "abc"

	//mark := "aaaaj"
	//str = ""



	index := strHasMark(mark, str)
	//
	//fmt.Println(index)
	//if index != -1 {
		fmt.Println(checkStrHasMark(mark, index, str))
		fmt.Println(mark, "-----模式串----")
		fmt.Println(searchModelStr(mark))
	//}
}