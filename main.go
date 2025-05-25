package main

import (
	"fmt"
	"math"
	"os"
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
	var sum3 = a + b // the type will be float64

	fmt.Println(sum1, sum2)
	fmt.Println("sum3:", sum3)
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

const x int8 = 12

func constPractice() {

	fmt.Println("-----constPractice")
	// x = x +1; will throw error even if it is declared outside of the function
	fmt.Println("x is", x)
	const y = 2
	// y = 3 will also fail

	const (
		z         = 4
		test bool = false
	)
	// z = true will also fail

	// a := 2
	// b := 3
	// const test = a + b will fail => a + b (value of type int) is not constant

	// var c int = 2
	// var d int = 2

	// const test2 = c + d will fail  => c + d (value of type int) is not constant

	const a = 2
	const v = 12
	const asd = a + v

	fmt.Printf("Valid assignment of %d and %d is %d \n", a, v, asd)

	const i = 2
	const f = 2.2
	const value = i + f

	fmt.Println(" the value is ", value)

}

func declarations() {
	fmt.Println("---------declarations")

	x := 2
	y := 2.2

	fmt.Printf("%d, %f \n", x, y)

	var i int8 = 127

	i = i + 4

	fmt.Println("outbound", i) // i = -125
}

func arrayDeclaration() {
	fmt.Println("---------array declaration")

	var noDeclarationInt [3]int
	fmt.Println("An integer array with default values: ", noDeclarationInt)

	var noDeclarationBool [3]bool
	fmt.Println("An bool array with default values: ", noDeclarationBool)

	var noDeclarationRune [3]rune
	fmt.Println("An rune array with default values: ", noDeclarationRune)

	literalDeclaration := [3]int{1, 2, 3}
	fmt.Println("An array declaration of integers with array literals: ", literalDeclaration)

	//accessing element of an array is similar to js
	// first index is 0
	fmt.Println("first index is 0, literalDeclaration first element is ", literalDeclaration[0])
	fmt.Println("second index is 1, literalDeclaration second element is ", literalDeclaration[1])
	fmt.Println("third index is 2, literalDeclaration third element is ", literalDeclaration[2])

	sparseArray := [12]int{1, 5: 3, 4, 10: 123, 11}

	fmt.Println("sparse array is ", sparseArray) // prints [1 0 0 0 0 3 4 0 0 0 123 11]

	// arrays are not immutable by default so we can not define as
	// const literalDeclaration = [3]int{1, 2, 3} -> it will be fails

	dotArrayDeclaration := [...]int{1, 2, 3}

	fmt.Println("dot array declaration: ", dotArrayDeclaration)

	isLiteralandDotEqual := literalDeclaration == dotArrayDeclaration

	// you can compare two arrays with ==
	fmt.Fprintf(os.Stdout, "compare of two same element array like literalDeclaration %v and dotArrayDeclaration %v is %t \n", []any{literalDeclaration, dotArrayDeclaration, isLiteralandDotEqual}...)
	differentElementPlacement := [...]int{1, 3, 2}
	isEqualNotOrder := differentElementPlacement == dotArrayDeclaration

	fmt.Fprintf(os.Stdout, "compare of two same element array like differentElementPlacement %v and dotArrayDeclaration %v is %t \n", []any{differentElementPlacement, dotArrayDeclaration, isEqualNotOrder}...)
}

func main() {
	runAlmostEqual()
	compareTwoDifferentType()
	declarations()
	compareTwoDifferentIntType()
	ifCheck()

	constPractice()

	arrayDeclaration()

}
