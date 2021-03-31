package dp

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	stairs := ClimbStairs(35)
	fmt.Print(stairs)
}

func TestIsSubsequence1(t *testing.T) {
	fmt.Println(IsSubsequence1("abcd", "sabc123abcd"))
}
