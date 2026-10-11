package main

import "fmt"

func takeTempirature() bool {
	var peop int

	_, err := fmt.Scan(&peop)
	if err != nil {
		fmt.Println("Invalid num of people")
		return false
	}

	var min_t int = 15
	var max_t int = 30

	for range peop {

		var op string
		_, err = fmt.Scan(&op)
		if err != nil || op != "<=" && op != ">=" {
			fmt.Println("Invalid operation")
			return false
		}

		var temp int
		_, err = fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Invalid temperature")
			return false
		}

		switch op {
		case ">=":
			min_t = max(min_t, temp)
		case "<=":
			max_t = min(max_t, temp)
		}

		if min_t <= max_t {
			fmt.Println(min_t)
		} else {
			fmt.Println(-1)
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
