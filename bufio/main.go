package main

import (
"bufio"
"fmt"
"strings"
)

func main() {

		data := "do rei me so fa la tee do"
		fmt.Printf("\nData:%s\n\n", data)

		scanner := bufio.NewScanner(strings.NewReader(data))

		for scanner.Scan() {
				fmt.Println(scanner.Text())
		}

}

