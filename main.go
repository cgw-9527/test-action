package main

import "fmt"

var version = "0.1.1"

func GetVersion() string {
	return version
}

func main() {
	fmt.Println(GetVersion())
}
