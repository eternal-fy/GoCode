package test

import "testing"

var in = []int{1, 2, 3, 4}
var out = []int{2, 3, 4, 5, 6}

func TestExample(t *testing.T) {
	if len(in) != len(out) {
		t.Errorf("the arrays of in and out is different!")
	}
	for i, _ := range in {
		if in[i]+1 != out[i] {
			t.Errorf("the index %d is error", i)
		}
	}
}
