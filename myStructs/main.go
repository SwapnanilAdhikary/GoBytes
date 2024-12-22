package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritence inb goland, no super parent, etc
	// struct == classes in java
	Swapnanil := User{"Swapnanil", "adhikaryswapnanil@gmail.com", true, 21}
	fmt.Println(Swapnanil)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
