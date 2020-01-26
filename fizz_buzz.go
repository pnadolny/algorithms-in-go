package algorithms_in_go

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//

func FizzBuzz(n int) {
	var sb strings.Builder
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			sb.WriteString("Fizz Buzz")
		} else if i%3 == 0 {
			sb.WriteString("Fizz")
		} else if i%5 == 0 {
			sb.WriteString("Buzz")
		} else {
			sb.WriteString(strconv.Itoa(i))
		}
		sb.WriteString(",")
	}
	finalString := strings.TrimSuffix(sb.String(), ",")
	fmt.Println(finalString)

}
