package main

import "fmt"

func main() {
	// var i = 0
	go func() {
		fmt.Println("11")
	}()
	go func() {
		fmt.Println("22")
	}()
	fmt.Println("00")
}
