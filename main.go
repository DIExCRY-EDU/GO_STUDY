package main

import (
	"fmt"
)

var a, b, c int

func main() {
	// message := []byte("asd")
	// fmt.Println(message)

	// var a byte = 97
	// fmt.Printf("%c\n", a)

	// var b rune = 'b'
	// fmt.Printf("%c\n", b)

	a, b, c := 1, 2, 3

	a, b = b, a

	a, _, c = 0, 2, 0
	fmt.Println(a, b, c)
}

func print() {
	fmt.Println(a, b, c)
}

// видео 45:00
