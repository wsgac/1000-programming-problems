package integer

import "testing"

func TestInteger15(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{123, 213},
		{542, 452},
		{921, 291},
	}
	for _, test := range tests {
		output := Integer15(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}

func TestInteger16(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{123, 132},
		{542, 524},
		{921, 912},
	}
	for _, test := range tests {
		output := Integer16(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}

func TestInteger17(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{3123, 1},
		{5542, 5},
		{1921, 9},
	}
	for _, test := range tests {
		output := Integer17(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}

func TestInteger18(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{3123, 3},
		{5542, 5},
		{1921, 1},
	}
	for _, test := range tests {
		output := Integer18(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}

func TestInteger24(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 0},
		{8, 1},
		{9, 2},
		{10, 3},
	}
	for _, test := range tests {
		output := Integer24(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}
func TestInteger25(t *testing.T) {
	tests := []struct {
		input, expected int64
	}{
		{1, 4},
		{2, 5},
		{3, 6},
		{4, 0},
		{5, 1},
		{6, 2},
		{7, 3},
		{8, 4},
		{9, 5},
		{10, 6},
	}
	for _, test := range tests {
		output := Integer25(test.input)
		if output != test.expected {
			t.Errorf("Got: %v Expected: %v", output, test.expected)
		}
	}
}
