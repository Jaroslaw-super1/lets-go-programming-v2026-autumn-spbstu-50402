package main

import "fmt"

func takeTempirature() bool {
	var peop int

	_, err := fmt.Scan(&peop)
	if err != nil {
		fmt.Println("Invalid num of people")

		return false
	}

	var minT, maxT int = 15, 30

	for range peop {
		var operation string
		if _, err := fmt.Scan(&operation); err != nil || operation != "<=" && operation != ">=" {
			fmt.Println("Invalid operation")

			return false
		}

		var temp int
		if _, err := fmt.Scan(&temp); err != nil {
			fmt.Println("Invalid temperature")

			return false
		}

		switch operation {
		case ">=":
			minT = max(minT, temp)
		case "<=":
			maxT = min(maxT, temp)
		}

		if minT <= maxT {
			fmt.Println(minT)
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
