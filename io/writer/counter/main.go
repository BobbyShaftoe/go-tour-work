package main

import (
	"fmt"
	_ "fmt"
)

type Counter int

func main() {
	fmt.Println("Write counter")

	var c Counter
	var d0 []byte
	var d1 []byte
	var d2 []byte

	c.Write(d0)
	c.Write(d1)
	c.Write(d2)

	fmt.Printf("Writer counter: %v\n", c)

}

func (c *Counter) Write(p []byte) (n int, err error) {
	*c = Counter(int(*c) + 1)
	fmt.Printf("Length of p: %v\n", len(p))
	return len(p), nil
}
