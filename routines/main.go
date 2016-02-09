package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("ROUTINE %d: %d\n", n, i)
		time.Sleep(time.Millisecond * 250)
	}
}
