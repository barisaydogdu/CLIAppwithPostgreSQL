package util

import "testing"

func TestSum(t *testing.T) {
	ss := Sum(65536, 65536)

	if ss != (65536 + 65536) {
		t.Error("not equal")
	}

}
