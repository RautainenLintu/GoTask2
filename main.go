package main

import "fmt"

func askForInput(name string) float64 {
	var number float64
	fmt.Printf("Enter your %s: ", name)
	fmt.Scan(&number)
	return number
}

func main() {
	income := askForInput("income")
	food := askForInput("food cost")
	transport := askForInput("transportation cost")
	accommodation := askForInput("accommodation cost")
	totalExpenses := food + transport + accommodation
	diff := income - totalExpenses
	fmt.Printf("Your surplus is: %f \n", diff)

	var expenses = [3]float64{food, transport, accommodation}
	var categories = [3]string{"food", "transport", "accommodation"}
	for index, item := range expenses {
		share := item / income
		if share >= 0.2 {
			fmt.Printf("Expenses are too high in the category: %s. They consist %f of your income.\n", categories[index], share)
		}
	}
}
