package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of All Languages",languages)
	fmt.Println("List of All Languages",languages["Js"])
	

	delete(languages,"Rb")
	fmt.Println("List of All Languages",languages)
	
	//loops are intersting in maps golang

	for key,value := range languages{
		fmt.Printf("For Key %v, value is %v\n",key,value)
	}

	
}