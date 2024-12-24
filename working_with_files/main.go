package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this needs to go in a file-test.com"

	file, err := os.Create("./mygo.txt")

	checkNillErr(err)
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is :", length)
	defer file.Close()
    readFile("./mygo.txt")
}

func readFile(filename string) {
	//the read data will be read in bytes format
	dataByte, err := ioutil.ReadFile(filename)

	checkNillErr(err)
	// fmt.Println("text data inside the file is \n", dataByte)
    fmt.Println("text data insde the file is \n",string(dataByte))

}


func checkNillErr(err error){
    if err != nil{
        panic(err)
    }
}