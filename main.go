package main
import "fmt"
import "math"


func calculateX(i, j , k, A, B, C float32) (x, y ,z float32){

	sinA, cosA  := math.Sin(A), math.Cos(A)
	sinB, cosB  := math.Sin(B), math.Cos(B)
	sinC, cosC  := math.Sin(C), math.Cos(C)

// Creating Rotational matricies for X
	x = j* sinA +sinB *cosC - k*cosA*sinB*cosC+
	j* cosA*sinC+k*sinA*sinC+j*cosB*cosC

// Creating Rotational matricies for Y
	y = j*cosA*cosC+k*sinA*cosC - j*sinA*sinB*sinC+k*cosA*sinB*sinC -
	i*cosB*sinC

// Creating Rotational matricies for Z
	z =k*cosA*cosB-j*sinA*cosB+j*sinB

	return
}

func main(){
	var cubeWidth float = 10
	var widrh, height int = 160, 40
	var zBuffer float32 = [160 *44]
	var buffer char = [160*44]
	var x, y, z float

	fmt.Println("Hello Donut Ascii")




}
