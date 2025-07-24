package main

import "fmt"
import "C"

func GetString() string {
	str := "Hello Gopher!"
	fmt.Println(str)
	return str
}

func main() {}
