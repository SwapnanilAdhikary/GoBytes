package main

import "fmt"

func main() {
	fmt.Println("welcome to functins in golang")
	greeter()
	result := adder(3,5)
	fmt.Println(result)

	proRes := proAdder(2,5,8,7)
	fmt.Println(proRes)
}

func adder(val1 int,val2 int) int{
	return val1+val2
}

func proAdder(values ...int) int{
	total := 0
	for _,val := range values{
		total+= val
	}
	return total
}

func greeter() {
	fmt.Println("Namstey from golang")
}
//
