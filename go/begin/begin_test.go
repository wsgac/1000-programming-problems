package begin

import (
	"math"
	"reflect"
	"testing"
)

func TestBegin1(t *testing.T) {
	tests := []struct {
		input, expected float64
	}{
		{1.0, 4.0},
		{10.0, 40.0},
		{2.0, 8.0},
		{123.0, 492.0},
	}
	for _, test := range tests {
		output := Begin1(test.input)
		if output != test.expected {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin2(t *testing.T) {
	tests := []struct {
		input, expected float64
	}{
		{1.0, 1.0},
		{10.0, 100.0},
		{2.0, 4.0},
		{123.0, 15129.0},
	}
	for _, test := range tests {
		output := Begin2(test.input)
		if output != test.expected {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin3(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.0, 2.0}, 6.0},
		{[]float64{10.0, 20.0}, 60.0},
		{[]float64{2.0, 7.0}, 18.0},
		{[]float64{123.0, 100.0}, 446.0},
	}
	for _, test := range tests {
		fun := reflect.ValueOf(Begin3)
		var args []reflect.Value
		for _, arg := range test.input {
			args = append(args, reflect.ValueOf(arg))
		}
		res := fun.Call(args)
		output := res[0].Interface().(float64)
		if output != test.expected {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin4(t *testing.T) {
	tests := []struct {
		input, expected float64
	}{
		{1.0, 3.14},
		{2.0, 6.28},
		{10.0, 31.4},
		{123.0, 386.22},
	}
	for _, test := range tests {
		output := Begin4(test.input)
		// Approximate equality due to float rounding
		if math.Abs(output-test.expected) > 1e-9 {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin5(t *testing.T) {
	tests := []struct {
		input    float64
		expected []float64
	}{
		{1.0, []float64{1.0, 6.0}},
		{2.0, []float64{8.0, 24.0}},
		{10.0, []float64{1000.0, 600.0}},
		{123.0, []float64{1860867.0, 90774.0}},
	}
	for _, test := range tests {
		vol, area := Begin5(test.input)
		// Approximate equality due to float rounding
		if vol != test.expected[0] || area != test.expected[1] {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{vol, area})
		}
	}
}

func TestBegin6(t *testing.T) {
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{1.0, 2.0, 3.0}, []float64{6.0, 22.0}},
		{[]float64{2.0, 4.0, 6.0}, []float64{48.0, 88.0}},
		{[]float64{10.0, 10.0, 20.0}, []float64{2000.0, 1000.0}},
		{[]float64{123.0, 128.0, 182.0}, []float64{2.865408e+06, 122852.0}},
	}
	for _, test := range tests {
		vol, area := Begin6(test.input[0], test.input[1], test.input[2])
		// Approximate equality due to float rounding
		if vol != test.expected[0] || area != test.expected[1] {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{vol, area})
		}
	}
}

func TestBegin7(t *testing.T) {
	tests := []struct {
		input    float64
		expected []float64
	}{
		{1.0, []float64{6.28, 3.14}},
		{2.0, []float64{12.56, 12.56}},
		{10.0, []float64{62.8, 314.0}},
		{123.0, []float64{772.44, 47505.06}},
	}
	for _, test := range tests {
		circ, area := Begin7(test.input)
		// Approximate equality due to float rounding
		if math.Abs(circ-test.expected[0]) > 1e-9 ||
			math.Abs(area-test.expected[1]) > 1e-9 {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{circ, area})
		}
	}
}

func TestBegin8(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.0, 4.0}, 2.5},
		{[]float64{2.0, 8.7}, 5.35},
		{[]float64{10.0, 1234.0}, 622.0},
		{[]float64{123.0, 128.0}, 125.5},
	}
	for _, test := range tests {
		output := Begin8(test.input[0], test.input[1])
		if output != test.expected {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin9(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1.0, 4.0}, 2.0},
		{[]float64{2.0, 8.7}, 4.171330722922842},
		{[]float64{10.0, 1234.0}, 111.08555261599052},
		{[]float64{123.0, 128.0}, 125.47509713086498},
	}
	for _, test := range tests {
		output := Begin9(test.input[0], test.input[1])
		if output != test.expected {
			t.Fatalf("Expected: %v Got: %v", test.expected, output)
		}
	}
}

func TestBegin22(t *testing.T) {
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{1.0, 4.0}, []float64{4.0, 1.0}},
		{[]float64{2.0, 8.7}, []float64{8.7, 2.0}},
		{[]float64{10.0, 1234.0}, []float64{1234.0, 10.0}},
		{[]float64{123.0, 128.0}, []float64{128.0, 123.0}},
	}
	for _, test := range tests {
		a, b := test.input[0], test.input[1]
		retA, retB := Begin22(&a, &b)
		if test.input[0] != b || test.input[1] != a ||
			test.input[0] != retB || test.input[1] != retA {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{a, b})
		}
	}
}

func TestBegin23(t *testing.T) {
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{1.0, 4.0, 43.0}, []float64{4.0, 43.0, 1.0}},
		{[]float64{2.0, 8.7, 832.1}, []float64{8.7, 832.1, 2.0}},
		{[]float64{10.0, 1234.0, 432.23}, []float64{1234.0, 432.23, 10.0}},
		{[]float64{123.0, 128.0, 52.234}, []float64{128.0, 52.234, 123.0}},
	}
	for _, test := range tests {
		a, b, c := test.input[0], test.input[1], test.input[2]
		retA, retB, retC := Begin23(&a, &b, &c)
		if test.input[0] != b || test.input[1] != c || test.input[2] != a ||
			test.input[0] != retB || test.input[1] != retC || test.input[2] != retA {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{a, b, c})
		}
	}
}

func TestBegin24(t *testing.T) {
	tests := []struct {
		input    []float64
		expected []float64
	}{
		{[]float64{1.0, 4.0, 43.0}, []float64{4.0, 43.0, 1.0}},
		{[]float64{2.0, 8.7, 832.1}, []float64{8.7, 832.1, 2.0}},
		{[]float64{10.0, 1234.0, 432.23}, []float64{1234.0, 432.23, 10.0}},
		{[]float64{123.0, 128.0, 52.234}, []float64{128.0, 52.234, 123.0}},
	}
	for _, test := range tests {
		a, b, c := test.input[0], test.input[1], test.input[2]
		retA, retB, retC := Begin24(&a, &b, &c)
		if test.input[0] != c || test.input[1] != a || test.input[2] != b ||
			test.input[0] != retC || test.input[1] != retA || test.input[2] != retB {
			t.Fatalf("Expected: %v Got: %v", test.expected, []float64{a, b, c})
		}
	}
}
