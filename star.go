package main

import "fmt"

func main() {
	num := 6
	Draw(num)
}

func Draw(num int) {
	for i := num; i > 1; i-- {
		for j := i; j > 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	for i := 3; i < num; i++ {
		for j := 0; j < i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
