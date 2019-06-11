package begin

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
