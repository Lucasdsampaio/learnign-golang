package main

import "fmt"

func main() {
	x := 1.2

	types := fmt.Sprintf("%T", x)
	fmt.Println(types)
	switch types {
	case "string":
		fmt.Println("stringui")
	case "int":
		fmt.Println("int")
	case "float64":
		fmt.Println("float")
	default:
		fmt.Println("another thing")
	}
}
