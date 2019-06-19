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

// Boolean16 checks if its argument is a two-digit even number
func Boolean16(n int64) bool {
	return n > 9 && n < 100 && n%2 == 0
}

// Boolean17 checks if its argument is a three-digit odd number
func Boolean17(n int64) bool {
	return n > 99 && n < 1000 && n%2 == 1
}

// Boolean18 checks if at least two of its arguments are equal
func Boolean18(a, b, c int64) bool {
	set := make(map[int64]bool)
	set[a] = true
	set[b] = true
	set[c] = true
	return len(set) < 3
}

// Boolean19 checks if at least two of its arguments are mutually opposite
func Boolean19(a, b, c int64) bool {
	return a == -b || a == -c || b == -c
}

// Boolean20 checks if all digits of a three digit number are different from each other
func Boolean20(n int64) bool {
	set := make(map[int64]bool)
	for i := 0; i < 3; i++ {
		set[n%10] = true
		n /= 10
	}
	return len(set) == 3
}

// Boolean21 checks if the digits of a three-digit number form an increasing sequence
func Boolean21(n int64) bool {
	return n%10 > n/10%10 && n/10%10 > n/100
}

// Boolean22 checks if the digits of a three-digit nubmer form a monotonic sequence
func Boolean22(n int64) bool {
	return (n%10 > n/10%10 && n/10%10 > n/100) ||
		(n%10 < n/10%10 && n/10%10 < n/100)
}

// Boolean23 checks if a four-digit number is a palindrome
func Boolean23(n int64) bool {
	str := string(n)
	return str[0] == str[3] && str[1] == str[2]
}

// Boolean24 checks if a quadratic equation has real roots
func Boolean24(a, b, c int64) bool {
	d := b*b - 4*a*c
	return d >= 0
}

// Boolean25 checks if point (x,y) lies in the 2nd quarter of the coordinate system
func Boolean25(x, y int64) bool {
	return x < 0 && y > 0
}

// Boolean26 checks if point (x,y) lies in the 4th quarter of the coordinate system
func Boolean26(x, y int64) bool {
	return x > 0 && y < 0
}

// Boolean27 checks if point (x,y) lies in the 3rd quarter of the coordinate system
func Boolean27(x, y int64) bool {
	return x < 0
}

// Boolean28 checks if point (x,y) lies in the 1st or 3rd quarter of the coordinate system
func Boolean28(x, y int64) bool {
	return (x > 0 && y > 0) || (x < 0 && y < 0)
}

// Boolean29 checks if point (x,y) lies inside a rectangle spanned by (x1,y1) and (x2,y2)
func Boolean29(x, y, x1, y1, x2, y2) bool {
	return x1 < x && x < x2 && y2 < y && y < y1
}

// Boolean30
func Boolean30(a, b, c int64) bool {

}
