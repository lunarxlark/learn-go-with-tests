package iteration

import (
	"fmt"
	"testing"
)

type testcase struct {
	str      string
	substr   string
	expected bool
}

func TestContain(t *testing.T) {

	testcases := []testcase{
		{str: "abc", substr: "a", expected: true},
		{str: "abc", substr: "ab", expected: true},
		{str: "abc", substr: "ac", expected: false},
		{str: "12345", substr: "123", expected: true},
		{str: "a1b2c3", substr: "a1b2c3", expected: true},
		{str: "a1b2c3", substr: "123", expected: false},
		{str: "a1b2c3", substr: "abc", expected: false},
	}
	assertMessage := func(t *testing.T, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			want := tc.expected
			got := Contain(tc.str, tc.substr)
			assertMessage(t, got, want)
		})
	}
}
