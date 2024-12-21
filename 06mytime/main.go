package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("Present Time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006  15:04:05 Monday")) //formatting date in golang

    createDate := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
    fmt.Println("Created Date is: ", createDate)
    fmt.Println("Created Date is: ", createDate.Format("01-02-2006  15:04:05 Monday"))
}
