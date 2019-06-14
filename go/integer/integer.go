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

// Integer7 returns the sum and product of n's digits
func Integer7(n int64) (int64, int64) {
	l := n / 10
	r := n % 10
	return l + r, l * r
}

// Integer8 returns number obtained from n by swapping its digits
func Integer8(n int64) int64 {
	return (n%10)*10 + (n / 10)
}

// Integer9 returns the first digit of a three-digit number
func Integer9(n int64) int64 {
	return n / 100
}

// Integer10 returns the third and second digits of a three-digit number
func Integer10(n int64) (int64, int64) {
	return n % 10, n / 10 % 10
}

// Integer11 returns the sum and products of digits of a three-digit number
func Integer11(n int64) (int64, int64) {
	var sum int64
	var prod int64
	for i := 0; i < 3; i++ {
		digit := int64(n % 10)
		n /= 10
		sum += digit
		prod *= digit
	}
	return sum, prod
}

// Integer12 inverts the order of digits in a three-digit number
func Integer12(n int64) int64 {
	var res int64
	for i := 0; i < 3; i++ {
		res = 10*res + n%10
		n /= 10
	}
	return res
}

// Integer13 rotates to the left digits of a three-digit number
func Integer13(n int64) int64 {
	return 10*n%100 + n/100
}

// Integer14 rotates to the right digits of a three-digit number
func Integer14(n int64) int64 {
	return 100*n%10 + n/10
}

// Integer15 swaps the first two digits in a three-digit number
func Integer15(n int64) int64 {
	d := make([]int64, 3)
	for i := 2; i >= 0; i-- {
		d[i] = n % 10
		n /= 10
	}
	d[0], d[1] = d[1], d[0]
	ret := int64(0)
	for i := 0; i < 3; i++ {
		ret = ret*10 + d[i]
	}
	return ret
}

// Integer16 swaps the last two digits in a three-digit number
func Integer16(n int64) int64 {
	d := make([]int64, 3)
	for i := 2; i >= 0; i-- {
		d[i] = n % 10
		n /= 10
	}
	d[2], d[1] = d[1], d[2]
	ret := int64(0)
	for i := 0; i < 3; i++ {
		ret = ret*10 + d[i]
	}
	return ret
}

// Integer17 returns the number of hundreds in a number > 999
func Integer17(n int64) int64 {
	return n / 100 % 10
}

// Integer18 returns the number of thousands in a number > 999
func Integer18(n int64) int64 {
	return n / 1000 % 10
}

// Integer19 return the number of full minutes in n seconds
func Integer19(n int64) int64 {
	return n / 60
}

// Integer20 return the number of full hours in n seconds
func Integer20(n int64) int64 {
	return n / 3600
}

// Integer21 return the number of seconds elapsed from the start of last full minute in n seconds
func Integer21(n int64) int64 {
	return n % 60
}

// Integer22 return the number of seconds elapsed from the start of last full hour in n seconds
func Integer22(n int64) int64 {
	return n % 3600
}

// Integer23 return the number of full minutes elapsed from the start of last full hour in n seconds
func Integer23(n int64) int64 {
	return n % 3600 / 60
}

// Integer24 returns the day of the week of the k-th day, assuming Jan 1 was Monday
func Integer24(k int64) int64 {
	return k % 7
}

// Integer25 returns the day of the week of the k-th day, assuming Jan 1 was Thursday
func Integer25(k int64) int64 {
	return (k + 3) % 7
}
