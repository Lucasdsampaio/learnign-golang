package main

import "fmt"

func main() {
	slice := []int{}
	for i := 1; i < 20; i += 2 {
		slice = append(slice, i)
	}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)
}
