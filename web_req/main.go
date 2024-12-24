package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://www.isro.gov.in/"

func main() {
	fmt.Println("web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("the type of response is %T\n", response)

	defer response.Body.Close() //callers responsibility
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length of the text is", length)
	fmt.Println("data saved in file")
}
