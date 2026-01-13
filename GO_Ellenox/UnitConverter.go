package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Ellenox!")
	fmt.Println("This is a Unit Converter Program")
	fmt.Println("1. Kilometers to Meters")
	fmt.Println("2. Meters to Kilometers")
	fmt.Println("3. Meters to Centimeters")
	fmt.Println("4. Centimeters to Meters")
	fmt.Println("5. Celsius to Fahrenheit")
	fmt.Println("6. Fahrenheit to Celsius")
	fmt.Println("7. Kilograms to Grams")
	fmt.Println("8. Grams to Kilograms")
	fmt.Println("9. Exit")
	fmt.Print("Enter a value (1-9): ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var km float64
		fmt.Println("Enter kilometers:")
		fmt.Scanf("%f", &km)
		fmt.Printf("Result: %.2f meters\n", km*1000)
	case 2:
		var m float64
		fmt.Println("Enter meters:")
		fmt.Scanf("%f", &m)
		fmt.Printf("Result: %.2f kilometers\n", m/1000)
	case 3:
		var m float64
		fmt.Println("Enter meters:")
		fmt.Scanf("%f", &m)
		fmt.Printf("Result: %.2f centimeters\n", m*100)
	case 4:
		var cm float64
		fmt.Println("Enter centimeters:")
		fmt.Scanf("%f", &cm)
		fmt.Printf("Result: %.2f meters\n", cm/100)
	case 5:
		var c float64
		fmt.Println("Enter Celsius:")
		fmt.Scanf("%f", &c)
		f := (c*9/5 + 32)
		fmt.Printf("Result: %.2f Fahrenheit\n", f)
	case 6:
		var f float64
		fmt.Println("Enter Fahrenheit:")
		fmt.Scanf("%f", &f)
		c := (f - 32) * 5 / 9
		fmt.Printf("Result: %.2f Celsius\n", c)
	case 7:
		var kg float64
		fmt.Println("Enter kilograms:")
		fmt.Scanf("%f", &kg)
		g := kg * 1000
		fmt.Printf("Result: %.2f grams\n", g)
	case 8:
		var g float64
		fmt.Println("Enter grams:")
		fmt.Scanf("%f", &g)
		k := g / 1000
		fmt.Printf("Result: %.2f kilograms\n", k)
	case 9:
	default:
		fmt.Println("Exiting the program.")
	}

	fmt.Println("Thank you for using the Unit Converter!")
}
