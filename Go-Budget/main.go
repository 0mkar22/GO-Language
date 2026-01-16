package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Expenses struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	Timestamp   string  `json:"timestamp"`
}

const file = "expenses.json"

func fileExists(filename string) (bool, error) {
	_, err := os.Lstat(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ReadExpensesFromFile(filename string) ([]Expenses, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// Parse JSON data into Expenses slice
	var expenses []Expenses
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func WriteExpensesToFile(filename string, expenses []Expenses) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fileExists, _ := fileExists(file)
	fmt.Println("Go-Budget Application")
	fmt.Println("Expense file exists:", fileExists)
	user_arguments := os.Args[1:]
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, list, summary, delete")
		os.Exit(1)
	}

	args := os.Args[2:]

	switch user_arguments[0] {
	case "add":
		HandleAdd(args)
	case "list":
		HandleList()
	case "summary":
		HandleSummary()
	case "delete":
		HandleDelete(args)
	default:
		fmt.Println("Unknown command. Available commands: add, list, summary, delete")

	}

}

func HandleAdd(args []string) {
	var expenses Expenses
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	addcmd.Parse(args)

	fmt.Println("Enter expense details:")
	fmt.Println("Amount:")
	fmt.Scanln(&expenses.Amount)
	fmt.Println("Description:")
	fmt.Scanln(&expenses.Description)
	fmt.Println("Category:")
	fmt.Scanln(&expenses.Category)
	expenses.Timestamp = time.Now().Format(time.RFC3339)

	fmt.Printf("%+v\n", expenses)

	if expenses.Amount == 0 || expenses.Description == "" || expenses.Category == "" {
		fmt.Println("Please provide all required fields: --amount, --desc, --cat")
		return
	}

	existing_expenses, err := ReadExpensesFromFile(file)
	if err != nil {
		expenses.ID = 1
	} else {
		expenses.ID = len(existing_expenses) + 1
	}

	all_expenses := append(existing_expenses, expenses)

	err = WriteExpensesToFile(file, all_expenses)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Expense added successfully.")
}

func HandleList() {
	expenses, err := ReadExpensesFromFile(file)
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}
	fmt.Println("List of Expenses:")
	for _, expense := range expenses {
		fmt.Printf("ID: %d, Description: %s, Amount: %.2f, Category: %s, Timestamp: %s\n",
			expense.ID, expense.Description, expense.Amount, expense.Category, expense.Timestamp)
	}
}

func HandleSummary() {
	expenses, err := ReadExpensesFromFile(file)
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}
	categoryTotals := make(map[string]float64)
	var overallTotal float64
	for _, expense := range expenses {
		categoryTotals[expense.Category] += expense.Amount
		overallTotal += expense.Amount
	}
	fmt.Println("Summary of Expenses:")
	for category, total := range categoryTotals {
		fmt.Printf("Category: %s, Total: %.2f\n", category, total)
	}
	fmt.Printf("Overall Total: %.2f\n", overallTotal)

	// ASCII Bar Chart
	totals := make(map[string]float64)

	for _, expense := range expenses {
		totals[expense.Category] += expense.Amount
	}
	max_val := 0.0
	for _, total := range totals {
		if total > max_val {
			max_val = total
		}
	}

	max_bar_length := 50
	block_char := "█" // Unicode U+2588

	fmt.Println("\nExpense Distribution:")
	for category, total := range totals {
		bar_length := int(total / max_val * float64(max_bar_length))
		bar := strings.Repeat(block_char, bar_length)
		fmt.Printf("%s: %s (%.2f)\n", category, bar, total)
	}
}

func HandleDelete(args []string) {
	deletecmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deletecmd.Int("id", 0, "ID of the expense to delete")
	deletecmd.Parse(args)
	flag.Parse()
	if *id == 0 {
		fmt.Println("Please provide the expense ID to delete using --id")
		return
	}
	expenses, err := ReadExpensesFromFile(file)
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}
	var updatedExpenses []Expenses
	found := false
	for _, expense := range expenses {
		if expense.ID == *id {
			found = true
			continue
		}
		updatedExpenses = append(updatedExpenses, expense)
	}
	if !found {
		fmt.Printf("Expense with ID %d not found.\n", *id)
		return
	}
	err = WriteExpensesToFile(file, updatedExpenses)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Expense deleted successfully.")
}
