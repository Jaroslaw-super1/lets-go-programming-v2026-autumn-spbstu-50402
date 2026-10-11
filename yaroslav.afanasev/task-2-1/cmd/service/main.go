package main

import "fmt"

func takeTempirature() bool {
	fmt.Println("sssss")
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
