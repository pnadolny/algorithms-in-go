package algorithms_in_go

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	tests := []struct {
		word string
		want string
	}{
		{"paul", "luap"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse = %v; want %v", got, tc.want)
			}
		})
	}

}

func TestReverseWithStringBuilder(t *testing.T) {

	ReverseWithStringBuilder("paul")

}
