package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	s := "GO"
	i := 42
	b := true

	fmt.Print("\n")

	fmt.Printf("Printing, %T (%q)! \n", s, s)
	fmt.Printf("Printing, %T (%d)! \n", i, i)
	fmt.Printf("Printing, %T (%t)! \n\n", b, b)

	fmt.Printf("Printing, %T (%v)! \n", s, s)
	fmt.Printf("Printing, %T (%v)! \n", i, i)
	fmt.Printf("Printing, %T (%v)! \n\n", b, b)

	fmt.Printf("Printing, %T (%#v)! \n", s, s)
	fmt.Printf("Printing, %T (%#v)! \n", i, i)
	fmt.Printf("Printing, %T (%#v)! \n\n", b, b)

	fmt.Printf("Printing, %T (%s)! \n", s, s)
}
