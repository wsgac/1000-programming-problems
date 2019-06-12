package integer

// Integer1 returns the number of full meters in l centimeters
func Integer1(l int64) int64 {
	return l / 100
}

// Integer2 returns the number of full tonnes in m kilograms
func Integer2(m int64) int64 {
	return m / 1000
}

// Integer3 returns the number of full kilobytes occupied by b bytes
func Integer3(b int64) int64 {
	return b / 1024
}

// Integer4 returns how many sections of length b will fit in length a
func Integer4(a, b int64) int64 {
	return a / b
}

// Integer5 returns minimal unused length remaining after filling a with b's
func Integer5(a, b int64) int64 {
	return a % b
}

// Integer6 returns the two digits of a two-digit number
func Integer6(n int64) (int64, int64) {
	return n / 10, n % 10
}

// Integer7 return the sum and product of n's digits
func Integer7(n int64) (int64, int64) {
	l := n / 10
	r := n % 10
	return l + r, l * r
}

// Integer8 return number obtained from n by swapping its digits
func Integer8(n int64) int64 {
	return (n%10)*10 + (n / 10)
}
