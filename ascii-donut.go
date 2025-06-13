package main

import "fmt"
import "time"
import "math"

func main() {

	const (
		width        = 160  // Width of the outputed donut frame
		height       = 40   // Height of the outputed donut frame
		thetaSpacing = 0.07 // angle step around the cross-section circle
		phiSpacing   = 0.02 // Scalling of the circle/tube of the donut
	)

	//Starting location of 3 dimentions
	//B setted 45 degree off
	A, B, C := 0.0, math.Pi/4, 0.0
	luminance := ".,-~:;=!*#$@"

	for {
		var zBuffer [width * height]float64
		var buffer [width * height]rune

		for i := range buffer {
			buffer[i] = ' '
			zBuffer[i] = 0
		}

		// calcX, calcY, calcZ := calculateX(x, y, z, A, B, C)

		// fmt.Println(calcX, calcY, calcZ)
		// fmt.Println("Hello Donut Ascii")

		for theta := 0.0; theta < 2*math.Pi; theta += thetaSpacing {
			for phi := 0.0; phi < 2*math.Pi; phi += phiSpacing {
				// circle position
				circleX := math.Cos(theta)
				circleY := math.Sin(theta)

				r1 := 2.0 // major radius
				r2 := 1.0 // minor radius

				// raw torus point in 3D
				x := (r2*math.Cos(phi) + r1) * circleX
				y := (r2*math.Cos(phi) + r1) * circleY
				z := r2 * math.Sin(phi)

				// rotate point using your rotational matrix
				xRot, yRot, zRot := calcRotMatrix(x, y, z, A, B, C)
				// xRot, yRot, zRot := rotateAxiasY(x, y, z, B)

				zRot += 5 // pull forward

				ooz := 1 / zRot
				xp := int(float64(width)/2 + 30*ooz*xRot)
				yp := int(float64(height)/2 - 15*ooz*yRot)

				// basic luminance calculation
				L := math.Cos(phi)*math.Cos(theta)*math.Sin(B) -
					math.Cos(A)*math.Cos(theta)*math.Sin(phi) -
					math.Sin(A)*math.Sin(theta) +
					math.Cos(B)*(math.Cos(A)*math.Sin(theta)-math.Cos(theta)*math.Sin(A)*math.Sin(phi))

				lumIndex := int(L * 8)
				if lumIndex < 0 {
					lumIndex = 0
				}
				if lumIndex > 11 {
					lumIndex = 11
				}

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

		// clear screen
		fmt.Print("\x1b[H") // ANSI escape: move cursor to top
		for i := range width * height {
			if i%width == 0 {
				fmt.Print("\n")
			}
			fmt.Printf("%c", buffer[i])
		}

		// rotate over time
		A += 0.04
		B += 0.02
		C += 0.01

		time.Sleep(30 * time.Millisecond)
	}
}

// TODO a triple check the matricies math behind it
func calcRotMatrix(i, j, k, A, B, C float64) (x, y, z float64) {

	sinA, cosA := math.Sin(A), math.Cos(A) // X
	sinB, cosB := math.Sin(B), math.Cos(B) // Y
	sinC, cosC := math.Sin(C), math.Cos(C) // Z

	// Rotation around Y-axis (B)
	x1 := i*cosB + k*sinB
	z1 := -i*sinB + k*cosB
	y1 := j

	// Rotation around X-axis (A)
	y2 := y1*cosA - z1*sinA
	z2 := y1*sinA + z1*cosA
	x2 := x1

	// Rotation around Z-axis (C)
	x = x2*cosC - y2*sinC
	y = x2*sinC + y2*cosC
	z = z2

	return
}

func rotateAxiasY(i, j, k, B float64) (x, y, z float64) {
	sinB, cosB := math.Sin(B), math.Cos(B)

	x = i*cosB + k*sinB
	y = j
	z = -i*sinB + k*cosB

	return
}
