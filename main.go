package main

import (
	"fmt"
	"math"
)

func almostEqual(a float64, b float64, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

func runAlmostEqual() {
	threshold := 1e-9
	fmt.Printf("The threshold is %f\n", threshold)
	fmt.Println(almostEqual(0.1+0.2, 0.3, -1))
}

func compareTwoDifferentType() {
	const a = 10
	var b float64 = 30.2

	var sum1 float64 = float64(a) + b
	var sum2 int = a + int(b)

	fmt.Println(sum1, sum2)
}

func compareTwoDifferentIntType() {
	var a int16 = 10000
	var b byte = 30

	var sum1 int16 = a + int16(b)
	var sum2 byte = byte(a) + b
	fmt.Println("compareTwoDifferentIntType", sum1, sum2)
}

func ifCheck() {
	var s string
	fmt.Println("s", s == "")
	if s == "" {
		fmt.Println("s is empty")
	}
}

func main() {
	runAlmostEqual()
	compareTwoDifferentType()
	compareTwoDifferentIntType()
	ifCheck()

	var (
		x int
		y int = 20
	)

	fmt.Println(x, y)

	// c := "I am a string"

	// c = 2
}
