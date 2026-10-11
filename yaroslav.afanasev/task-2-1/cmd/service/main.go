package main

import "fmt"

func takeTempirature() bool {
	var peop int

	_, err := fmt.Scan(&peop)
	if err != nil {
		fmt.Println("Invalid num of people")
		return false
	}

	var min int = 10
	var max int = 30

	for range peop {
		var op string

		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Println("Invalid operation")
			return false
		}
	}

	return true
}

func main() {
	var depart int

	_, err := fmt.Scan(&depart)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	for range depart {
		if !takeTempirature() {
			return
		}
	}
}
