package main

import (
	"fmt"
	"strconv"
)

// import "os"
// import "errors"

func askForInput(name string) int {
	var number int
	var input string
	fmt.Printf("Enter your %s: ", name)
	fmt.Scan(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return -1
	}
	return number
}

func main() {
	income := askForInput("income")
	if income < 0 {
		fmt.Print("Invalid income. Restart the program and try again\n")
		return
	}
	food := askForInput("food cost")
	transport := askForInput("transportation cost")
	accommodation := askForInput("accommodation cost")
	var expenses = [3]int{food, transport, accommodation}
	for _, expense := range expenses {
		if expense < 0 {
			fmt.Print("Invalid expense. Restart the program and try again\n")
			return
		}
	}
	totalExpenses := food + transport + accommodation
	diff := income - totalExpenses
	if diff < 0 {
		fmt.Print("Expenses are greater than income. Check the numbers, restart the program and try again\n")
		return
	}
	fmt.Printf("Your surplus is: %d \n", diff)
	var categories = [3]string{"food", "transport", "accommodation"}
	for index, item := range expenses {
		share := float64(item) / float64(income)
		if share >= 0.2 {
			fmt.Printf("Expenses are too high in the category: %s. They consist %d percent of your income.\n", categories[index], int(share*100))
		}
	}
}
