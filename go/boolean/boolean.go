package boolean

// Boolean1 checks if its argument is a positive number
func Boolean1(a int64) bool {
	return a > 0
}

// Boolean2 checks if its argument is odd
func Boolean2(a int64) bool {
	return a%2 == 1
}

// Boolean3 checks if its argument is even
func Boolean3(a int64) bool {
	return a%2 == 0
}

// Boolean4 checks if a > 2 and b <= 3
func Boolean4(a, b int64) bool {
	return a > 2 && b <= 3
}

// Boolean5 checks whether a >= 0 or b < -2
func Boolean5(a, b int64) bool {
	return a >= 0 || b < -2
}

// Boolean6 checks if a < b < c
func Boolean6(a, b, c int64) bool {
	return a < b && b < c
}

// Boolean7 checks if b falls between a and c
func Boolean7(a, b, c int64) bool {
	return a <= b && b <= c
}

// Boolean8 checks if both its arguments are odd
func Boolean8(a, b int64) bool {
	return a%2 == 1 && b%2 == 1
}

// Boolean9 checks if at least one of its arguments is odd
func Boolean9(a, b int64) bool {
	return a%2 == 1 || b%2 == 1
}

// Boolean10 checks if exactly one of its arguments is odd
func Boolean10(a, b int64) bool {
	return (a%2 == 1) != (b%2 == 1)
}

// Boolean11 checks if its arguments have the same parity
func Boolean11(a, b int64) bool {
	return (a^b)%2 == 0
}

// Boolean12 checks if all its arguments are positive
func Boolean12(a, b, c int64) bool {
	return a > 0 && b > 0 && c > 0
}

// Boolean13 checks if at least one of its arguments is positive
func Boolean13(a, b, c int64) bool {
	return a > 0 || b > 0 || c > 0
}

// Boolean14 checks if exactly one of its arguments is positive
func Boolean14(a, b, c int64) bool {
	cnt := 0
	if a > 0 {
		cnt++
	}
	if b > 0 {
		cnt++
	}
	if c > 0 {
		cnt++
	}
	return cnt == 1
}

// Boolean15 checks if exactly two of its arguments are positive
func Boolean15(a, b, c int64) bool {
	cnt := 0
	if a > 0 {
		cnt++
	}
	if b > 0 {
		cnt++
	}
	if c > 0 {
		cnt++
	}
	return cnt == 2
}
