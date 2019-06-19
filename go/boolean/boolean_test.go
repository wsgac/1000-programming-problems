package boolean

import "testing"

func TestBoolean11(t *testing.T) {
	tests := []struct {
		input  []int64
		output bool
	}{
		{[]int64{10, 12}, true},
		{[]int64{11, 13}, true},
		{[]int64{10, 13}, false},
		{[]int64{11, 12}, false},
	}
	for _, test := range tests {
		output := Boolean11(test.input[0], test.input[1])
		if output != test.output {
			t.Errorf("Expected: %v Got: %v", test.output, output)
		}
	}
}

func TestBoolean18(t *testing.T) {
	tests := []struct {
		input  []int64
		output bool
	}{
		{[]int64{1, 2, 3}, false},
		{[]int64{1, 1, 2}, true},
		{[]int64{1, 2, 2}, true},
		{[]int64{1, 1, 1}, true},
	}
	for _, test := range tests {
		output := Boolean18(test.input[0], test.input[1], test.input[2])
		if output != test.output {
			t.Errorf("Expected: %v Got: %v", test.output, output)
		}
	}
}

func TestBoolean20(t *testing.T) {
	tests := []struct {
		input  int64
		output bool
	}{
		{123, true},
		{112, false},
		{122, false},
		{111, false},
	}
	for _, test := range tests {
		output := Boolean20(test.input)
		if output != test.output {
			t.Errorf("Expected: %v Got: %v", test.output, output)
		}
	}
}
