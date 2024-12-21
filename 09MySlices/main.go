package main

import "fmt"

func main() {
	fmt.Println("Welcome to Video on SLices")

	var fruitList = []string{"Apple", "Banana"}
	fmt.Println("Type of Fruitlist is %T", fruitList)
	fruitList = append(fruitList, "Manga")
	fmt.Println(fruitList)
}
