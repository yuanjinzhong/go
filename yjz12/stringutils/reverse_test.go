package xxx

import (
	"fmt"
	"testing"
	"go_respority/go_utils"
)

func TestReverse(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		fmt.Print(got)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestStringUtil(t *testing.T) {



}
