package begin

import "math"

// Begin1 calculates the perimeter of a square based on its side
func Begin1(a float64) float64 {
	return a * 4.0
}

// Begin2 calculates the surface area of a square based on its side
func Begin2(a float64) float64 {
	return a * a
}

// Begin3 calculates the perimeter of a rectangle based on its sides
func Begin3(a, b float64) float64 {
	return 2.0 * (a + b)
}

// Begin4 calculates the approximate circumference of a circle based on its diameter
func Begin4(d float64) float64 {
	return 3.14 * d
}

// Begin5 calculates the volume and surface area of a cube based on its side
func Begin5(a float64) (float64, float64) {
	face := a * a
	return face * a, 6.0 * face
}

// Begin6 calculates the volume and surface area of a parallelpiped based on its sides
func Begin6(a, b, c float64) (float64, float64) {
	face := a * b
	return face * c, 2.0 * (face + b*c + a*c)
}

// Begin7 calculates the circumference and surface area of a circle based on its radius
func Begin7(r float64) (float64, float64) {
	// Calculate once to save one multiplication
	PiR := 3.14 * r
	return 2 * PiR, PiR * r
}

// Begin8 calculates the arithmetic mean of its arguments
func Begin8(a, b float64) float64 {
	return (a + b) / 2.0
}

// Begin9 calculates the geometrical mean of its arguments
func Begin9(a, b float64) float64 {
	return math.Sqrt(a * b)
}

// Begin10 calculates the sum, difference, product and quotient of squares of two numbers
func Begin10(a, b float64) (float64, float64, float64, float64) {
	aSq := a * a
	bSq := b * b
	return aSq + bSq, aSq - bSq, aSq * bSq, aSq / bSq
}

// Begin11 calculates the sum, difference, product and quotient of modules of two numbers
func Begin11(a, b float64) (float64, float64, float64, float64) {
	aAbs := math.Abs(a)
	bAbs := math.Abs(b)
	return aAbs + bAbs, aAbs - bAbs, aAbs * bAbs, aAbs / bAbs
}

// Begin12 calculates the hypothenuse and perimeter of a right triangle based on its two catheti (legs)
func Begin12(a, b float64) (float64, float64) {
	c := math.Sqrt(a*a + b*b)
	return c, a + b + c
}

// Begin13 calculates the surface areas of two circles and that of the annulus spanned between the tow radii
func Begin13(r1, r2 float64) (float64, float64, float64) {
	s1 := 3.14 * r1 * r1
	s2 := 3.14 * r2 * r2
	return s1, s2, s1 - s2
}

// Begin14 calculates the radius and surface area of a circle based on its circumference
func Begin14(l float64) (float64, float64) {
	r := l / (2 * 3.14)
	return r, 3.14 * r * r
}

// Begin15 calculates diameter and circumference of a circle based on its surface area
func Begin15(s float64) (float64, float64) {
	r := math.Sqrt(s / 3.14)
	return 2 * r, 2 * 3.14 * r
}

// Begin16 calculates the length of the section |x1, x2|
func Begin16(x1, x2 float64) float64 {
	return math.Abs(x2 - x1)
}

// Begin17 calculates |AC|, |BC| and their sum
func Begin17(a, b, c float64) (float64, float64, float64) {
	ac := math.Abs(c - a)
	bc := math.Abs(c - b)
	return ac, bc, ac + bc
}

// Begin18 calculates the lengths of |AC| and |BC|
func Begin18(a, b, c float64) (float64, float64) {
	return math.Abs(c - a), math.Abs(c - b)
}

// Point2d is a struct representing a two-dimensional point
type Point2d struct {
	x, y float64
}

// Point3d is a struct representing a three-dimensional point
type Point3d struct {
	Point2d
	z float64
}

// Begin19 calculates perimeter and surface area of a rectangle based on its two opposite vertices
func Begin19(p1, p2 Point2d) (float64, float64) {
	a := math.Abs(p1.x - p2.x)
	b := math.Abs(p1.y - p2.y)
	return 2 * (a + b), a * b
}

// Begin20 calculates the distance between two points
func Begin20(p1, p2 Point2d) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2.0) + math.Pow(p1.y-p2.y, 2.0))
}

// Begin21 calculates perimeter and area of a triangle using Hero's formula
func Begin21(p1, p2, p3 Point2d) (float64, float64) {
	a := Begin20(p1, p2)
	b := Begin20(p2, p3)
	c := Begin20(p3, p1)
	p := (a + b + c) / 2.0
	return 2 * p, math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

// Begin22 swaps its two arguments and returns their new values
func Begin22(a, b *float64) (float64, float64) {
	*a, *b = *b, *a
	return *a, *b
}

// Begin23 rotates the values of a, b and c to the right
func Begin23(a, b, c *float64) (float64, float64, float64) {
	*a, *b, *c = *c, *a, *b
	return *a, *b, *c
}

// Begin24 rotates the values of a, b and c to the left
func Begin24(a, b, c *float64) (float64, float64, float64) {
	*a, *b, *c = *b, *c, *a
	return *a, *b, *c
}

// Begin25 calculates the appropriate polynomial at x
func Begin25(x float64) float64 {
	x2 := x * x
	return 3*x2*x2*x2 - 6*x2 - 7
}

// Begin26 calculates the appropriate polynomial in a shifted point
func Begin26(x float64) float64 {
	shifted3 := (x - 3) * (x - 3) * (x - 3)
	return 4*shifted3*shifted3 - 7*shifted3 + 2
}

// Begin27 calculates a^8 using 3 multiplications
func Begin27(a float64) float64 {
	tmp := a * a
	tmp *= tmp
	return tmp * tmp
}

// Begin28 calculates a^15 using 5 multiplications and 2 tmp variables
func Begin28(a float64) float64 {
	tmp2 := a * a
	tmp3 := tmp2 * a
	tmp2 *= tmp2 * tmp3
	tmp3 *= tmp2 * tmp2
	return tmp2 * tmp3
}

// Begin29 converts degrees to radians
func Begin29(a float64) float64 {
	return a / 180.0 * 3.14
}

// Begin30 converts radians to degrees
func Begin30(a float64) float64 {
	return a / 3.14 * 180.0
}
