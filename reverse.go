package algorithms_in_go

import (
	"fmt"
	"strings"
)

func Reverse(word string) string {
	var res = ""
	for i:= len(word) -1; i>=0; i-- {
		res = res + string(word[i])
	}
	fmt.Print(res)
	return res

}

func ReverseWithStringBuilder(word string) string {
	var s strings.Builder
	for i:= len(word) -1; i>=0; i-- {
		s.WriteByte(word[i])
	}
	fmt.Println(s.String())
	return s.String()

}