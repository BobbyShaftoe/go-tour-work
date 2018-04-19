package main

import (
		"fmt"
		_ "fmt"
)

type Counter struct {
		counter int
}

func main() {
		fmt.Println("Bytes counter")

		var c Counter
		var d0 = []byte("Hello, world!")
		var d1 []byte
		var d2 []byte

		c.Write(d0)
		c.Write(d1)
		c.Write([]byte{0, 5, 9, 10})
		c.Write(d2)
		c.Write([]byte{11, 1, 2})

		fmt.Printf("Bytes counter: %v\n", c.counter)

}

func (c *Counter) Write(p []byte) (n int, err error) {
		l := len(p)
		c.counter = c.counter + l
		return l, nil
}
