package main

import(
	"fmt"
	"math"
)

func sum(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mul(x int, y int) int {
	return x*y
}

func circle(R float32) float32 {
	A := math.Pi*(R*R)
	return A
}

func main(){
	fmt.Println("Sum: ", sum(1, 2))
	fmt.Println("Sub: ", sub(1, 2))
	fmt.Println("Mul: ", mul(1, 2))	
	fmt.Println("Area Circle: ", circle(5.2))
}