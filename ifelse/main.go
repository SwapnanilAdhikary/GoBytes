package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")
	var result string
	loginCount := 4
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Numver is Odd")
	}

	if num := 3; num < 10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less than 10")
	}
	
}
