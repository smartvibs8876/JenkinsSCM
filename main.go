package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Jenkins CICD", i+1)
	}
	fmt.Println("Testing VSC git")
}
