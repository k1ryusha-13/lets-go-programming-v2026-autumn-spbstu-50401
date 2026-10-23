package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departmentCount int

	fmt.Scan(&departmentCount)

	for department := 0; department < departmentCount; department++ {
		var employeeCount int
		fmt.Scan(&employeeCount)

		lowerBound := minTemperature
		upperBound := maxTemperature

		for employee := 0; employee < employeeCount; employee++ {
			var operation string
			var temperature int

			fmt.Scan(&operation, &temperature)

			switch operation {
			case ">=":
				if temperature > lowerBound {
					lowerBound = temperature
				}
			case "<=":
				if temperature < upperBound {
					upperBound = temperature
				}
			}

			if lowerBound <= upperBound {
				fmt.Println(lowerBound)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
