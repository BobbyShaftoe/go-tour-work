package main

import (
	"fmt"
	_ "fmt"
	"math"
	"math/cmplx"
	_ "math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	radius float64    = 10.55
)

const (
	Pi     = 3.141592
	Big    = 1 << 100
	Small  = Big >> 99
	Little = 2 << 15
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Convert types
	var x, y int = 3, 101
	var g float64 = math.Sqrt(float64(x*x + y*y))
	var r uint = uint(g)

	fmt.Println(x, y, r)

	fmt.Println(Pi * radius * radius)

	fmt.Println(math.Sqrt(Little))

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
