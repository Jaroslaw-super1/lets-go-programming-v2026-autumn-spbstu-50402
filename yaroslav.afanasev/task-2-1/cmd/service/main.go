package main

import "fmt"

const (
	initialMinTemperature = 15
	initialMaxTemperature = 30
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange() TemperatureRange {
	return TemperatureRange{
		min: initialMinTemperature,
		max: initialMaxTemperature,
	}
}

func (r *TemperatureRange) Apply(operation string, temperature int) error {
	switch operation {
	case ">=":
		r.min = max(r.min, temperature)
	case "<=":
		r.max = min(r.max, temperature)
	default:
		return fmt.Errorf("invalid operation %q: expected >= or <=", operation)
	}

	return nil
}

func (r TemperatureRange) ComfortableMin() (int, bool) {
	if r.min <= r.max {
		return r.min, true
	}

	return -1, false
}

func readInt(what string) (int, error) {
	var value int
	if _, err := fmt.Scan(&value); err != nil {
		return 0, fmt.Errorf("read %s: %w", what, err)
	}

	return value, nil
}

func readOperation() (string, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil {
		return "", fmt.Errorf("read operation: %w", err)
	}

	return operation, nil
}

func processDepartment() error {
	people, err := readInt("number of people")
	if err != nil {
		return err
	}

	temperatureRange := NewTemperatureRange()

	for i := 0; i < people; i++ {
		operation, err := readOperation()
		if err != nil {
			return err
		}

		temperature, err := readInt("temperature")
		if err != nil {
			return err
		}

		if err := temperatureRange.Apply(operation, temperature); err != nil {
			return err
		}

		if minTemperature, ok := temperatureRange.ComfortableMin(); ok {
			fmt.Println(minTemperature)
		} else {
			fmt.Println(-1)
		}
	}

	return nil
}

func main() {
	departments, err := readInt("number of departments")
	if err != nil {
		fmt.Println(err)

		return
	}

	for i := 0; i < departments; i++ {
		if err := processDepartment(); err != nil {
			fmt.Println(err)

			return
		}
	}
}
