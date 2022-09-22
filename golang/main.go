package main

import (
	"fmt"

	"./lifo"
)

func main() {

	testLifo()
}

func testLifo() {

	l := lifo.Lifo{}

	fmt.Printf("New stack: %s", l.PrintStack())

	l.Append("My first element")
	fmt.Printf("\nStack:\n%sTop = %s\n", l.PrintStack(), l.PrintTop())

	l.Append(2)
	fmt.Printf("\nStack:\n%sTop = %s\n", l.PrintStack(), l.PrintTop())

	l.Append(3)
	fmt.Printf("\nStack:\n%sTop = %s\n", l.PrintStack(), l.PrintTop())
}
