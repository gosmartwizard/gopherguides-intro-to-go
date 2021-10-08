package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	s := "GO"
	i := 42
	varBoolean := true

	fmt.Print("\n")

	fmt.Printf("Printing, %T (%q)! \n", varString, varString)
	fmt.Printf("Printing, %T (%d)! \n", varInteger, varInteger)
	fmt.Printf("Printing, %T (%t)! \n\n", varBoolean, varBoolean)

	fmt.Printf("Printing, %T (%v)! \n", varString, varString)
	fmt.Printf("Printing, %T (%v)! \n", varInteger, varInteger)
	fmt.Printf("Printing, %T (%v)! \n\n", varBoolean, varBoolean)

	fmt.Printf("Printing, %T (%#v)! \n", varString, varString)
	fmt.Printf("Printing, %T (%#v)! \n", varInteger, varInteger)
	fmt.Printf("Printing, %T (%#v)! \n\n", varBoolean, varBoolean)

	fmt.Printf("Printing, %T (%s)! \n", varString, varString)
}
