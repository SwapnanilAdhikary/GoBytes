package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Our Pizz:\n")
	// comma ok // err error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
