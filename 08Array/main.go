package main

import "fmt"

func main() {
	fmt.Println("Welcom to array in golang")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Lichi"
	fruitList[3] = "Mango"
	fruitList[4] = "lemon"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("length of arry is :",len(fruitList))
}
//
