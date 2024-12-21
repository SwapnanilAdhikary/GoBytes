package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcomme to our Pizza App")
	fmt.Println("Please rate our Pizza between 1 and 5: ")	
    reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
    fmt.Println("Thanks for rating Our Pizza: ", input)
    fmt.Printf("Type of rating is %T ", input) 
	numrating, err  := strconv.ParseFloat(strings.TrimSpace(input),64)
    if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("Added one to your rating:",numrating+1)
	}
}