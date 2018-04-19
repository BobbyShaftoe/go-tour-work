package main

import (
	"errors"
	"fmt"
	_ "fmt"
	"log"
)

type Person struct {
	name string
	age  uint8
	ssn  string
}

func main() {
	fmt.Println("Person reader")

	bob := Person{"Bobby Jones", 35, "100-01-1101"}
	mary := Person{"Mary-Anne Smith", 83, "831-01-8101"}
	jane := Person{"Jane Doe", 17, "021-23-2376"}

	var buf = make([]byte, 200)

	offset, n := 0, 0
	var err error

	n, err = bob.Read(buf[offset:])
	if err != nil {
		fmt.Println(buf)
		log.Fatalf("n: %v, err %v\n", n, err)
	}
	offset += n
	n, err = mary.Read(buf[offset:])
	if err != nil {
		fmt.Println(buf)
		log.Fatalf("n: %v, err %v\n", n, err)
	}
	offset += n
	n, err = jane.Read(buf[offset:])
	if err != nil {
		fmt.Println(buf)
		log.Fatalf("n: %v, err %v\n", n, err)
	}
	offset += n

	fmt.Println(buf)
}

/*
	Format: [
				<name-len>: byte
				<name>: []byte
				<age>: byte
				<ssn>: []byte
	]
*/

func (r Person) Read(p []byte) (n int, err error) {

	var l = len(r.name)
	var buf = make([]byte, l+1)
	buf[0] = byte(l)

	// encode r.name
	for i := 0; i < l; i++ {
		buf[i+1] = r.name[i]
	}

	// encode r.age
	buf = append(buf, byte(r.age))

	// encode r.ssn
	l = len(r.ssn)

	for i := 0; i < l; i++ {
		buf = append(buf, r.ssn[i])
	}

	n = copy(p, buf)

	if n < len(buf) {
		s := fmt.Sprintf("Buffer not big enough, required: %v, provided: %v", len(buf), len(p))
		return n, errors.New(s)
	}

	return n, nil
}

// String func
func (r Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v, ssn: %v)",
		r.name, r.age, r.ssn)
}
