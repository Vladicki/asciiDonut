package main

import "fmt"
import "math"

func main() {

	const (
		width        = 160
		height       = 40
		thetaSpacing = 0.07 // angle step around the cross-section circle
		phiSpacing   = 0.02
	)

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

	//TODO remove it and add the rotational matrx
	sinA, cosA := math.Sin(A), math.Cos(A)
	sinB, cosB := math.Sin(B), math.Cos(B)

	fmt.Println(calcX, calcY, calcZ)
	fmt.Println("Hello Donut Ascii")

	for theta := 0.0; theta < 2*math.Pi; theta += thetaSpacing {
		for phi := 0.0; phi < 2*math.Pi; phi += phiSpacing {
			// 3D coordinates of the surface of the torus
			circleX := math.Cos(theta)
			circleY := math.Sin(theta)

			// torus radius setup
			r2 := 1.0
			r1 := 2.0

			// 3D coordinates after rotating the torus
			x := (r2*math.Cos(phi) + r1) * circleX
			y := (r2*math.Cos(phi) + r1) * circleY
			z := r2 * math.Sin(phi)
			// 3D rotation (only A and B for now)
			xRot := cosB*x + sinA*sinB*y + cosA*sinB*z
			yRot := cosA*y - sinA*z
			zRot := -sinB*x + sinA*cosB*y + cosA*cosB*z + 5 // +5: pull donut away

			ooz := 1 / zRot
			xp := int(float64(width)/2 + 30*ooz*xRot)
			yp := int(float64(height)/2 - 15*ooz*yRot)

			// simple luminance / brightness value
			L := math.Cos(phi)*math.Cos(theta)*sinB -
				cosA*math.Cos(theta)*math.Sin(phi) -
				sinA*math.Sin(theta) +
				cosB*(cosA*math.Sin(theta)-math.Cos(theta)*sinA*math.Sin(phi))

			lumIndex := int(L * 8)
			if lumIndex < 0 {
				lumIndex = 0
			}
			if lumIndex > 11 {
				lumIndex = 11
			}
			luminance := ".,-~:;=!*#$@"
			char := rune(luminance[lumIndex])

			idx := xp + yp*width
			if idx >= 0 && idx < width*height {
				if ooz > zBuffer[idx] {
					zBuffer[idx] = ooz
					buffer[idx] = char
				}
			}
		}
	}

	// Print buffer
	for i := range width * height {
		if i%width == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%c", buffer[i])
	}
}

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
