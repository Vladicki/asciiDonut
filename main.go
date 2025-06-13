package main

import "fmt"
import "math"

func calculateX(i, j, k, A, B, C float64) (x, y, z float64) {

	sinA, cosA := math.Sin(A), math.Cos(A)
	sinB, cosB := math.Sin(B), math.Cos(B)
	sinC, cosC := math.Sin(C), math.Cos(C)

	// Creating Rotational matricies for X
	x = j*sinA*sinB*cosC - k*cosA*sinB*cosC +
		j*cosA*sinC + k*sinA*sinC + j*cosB*cosC

	// Creating Rotational matricies for Y
	y = j*cosA*cosC + k*sinA*cosC - j*sinA*sinB*sinC + k*cosA*sinB*sinC -
		i*cosB*sinC

		// Creating Rotational matricies for Z
	z = k*cosA*cosB - j*sinA*cosB + j*sinB

	return
}

func main() {

	const (
		width  = 160
		height = 40
	)
	// var cubeWidth float64 = 10
	// var width, height int = 160, 40
	var zBuffer [width * height]float64
	var buffer [width * height]rune

	for i := range buffer {
		buffer[i] = ' '
		zBuffer[i] = 0
	}
	fmt.Println(buffer)

	x, y, z := 1.0, 0.0, 0.0

	A := math.Pi / 4 // degree
	B := math.Pi / 4 // degree
	C := math.Pi / 4 // degree
	calcX, calcY, calcZ := calculateX(x, y, z, A, B, C)

	fmt.Println(calcX, calcY, calcZ)
	fmt.Println("Hello Donut Ascii")

}
