package main

import (
	"fmt"
	_ "fmt"
	"math"
	"time"
)

func pow(x, n, lim float64) float64 {
	// If statement can have some execute statement before condition
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("Value %g to power of %g, hit limit: %g\n", x, n, lim)
	}
	return lim
}

func switchit() string {
	// case statement
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Far away")
	}
	return "done"
}

func deferit() string {
	// defer statement, stacked defers are last in first out
	defer fmt.Println("world")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
	return "done"
}

func main() {
	sum := 10
	for i := 0; i < 10; i++ {
		sum += i - 3
	}
	fmt.Println(sum)

	// Don't need init or post in loop declaration
	for sum < 1000 {
		sum += sum / 2
	}
	fmt.Println(sum)

	// Don't need semicolons
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// endless loop - like while true statement
	for {
		sum = sum / 2
		fmt.Println(sum)
		if sum < 1 {
			break
		}
	}

	fmt.Println(
		pow(3, 3, 30),
		pow(9, 3, 80),
	)

	switchit()
	deferit()

}




// import (
// "fmt"
// "math"
// )
//
// var x = float64(100199191111)
// var y = float64(x-1)
//
// func Sqrt(x float64) float64 {
// 	z := y
// 	q := z
// 	var diff float64
//
// 	for count := 1; count < 30; count++ {
// 		z = (z + x/z) / 2
// 		fmt.Println("difference:", diff)
// 		diff = math.Abs(q - z)
//
//
// 		q = z
// 		//	fmt.Println("next guess:", q)
// 	}
// 	return z
// }
//
// func main() {
// 	fmt.Println("Calculating root for:", x)
// 	fmt.Println("Initial guess is:", y)
// 	fmt.Println(Sqrt(x))
// 	fmt.Println(math.Sqrt(x))
// }