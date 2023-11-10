package main

import (
	"fmt"
	"time"
)

func main() {
	cafe := GetCafeInstance()
	fmt.Println("Welcome to the cafe!")
	fmt.Println("Can you tell me your name?")
	var name string
	fmt.Scanln(&name)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Place an order")
		fmt.Println("2. View available coffee drinks")
		fmt.Println("3. Apply decorator to a coffee")
		fmt.Println("4. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Place an order - Choose a coffee type (espresso, latte, cappuccino, americano, mocha, macchiato):")
			var coffeeChoice string
			fmt.Scanln(&coffeeChoice)

			coffeeDrink, err := getCoffeeDrink(coffeeChoice)
			if err != nil {
				fmt.Println(err)
			} else {
				order := Order{
					Coffee: Coffee{Name: coffeeDrink.getName()},
					Status: "Preparing",
				}
				coffeeOrder := CoffeeOrder{Order: order}

				// Creating observers (customers)
				customer4 := Customer{Name: name}
				customer1 := Customer{Name: "Raiymbek"}
				customer2 := Customer{Name: "Miras"}
				customer3 := Customer{Name: "Aisultan"}
				time.Sleep(3 * time.Second)
				coffeeOrder.ChangeStatus("Ready")
				coffeeOrder.RegisterObserver(customer4)
				coffeeOrder.RegisterObserver(customer1)
				coffeeOrder.RegisterObserver(customer2)
				coffeeOrder.RegisterObserver(customer3)
				coffeeOrder.ChangeStatus("Ready")
				cafe.Orders = append(cafe.Orders, coffeeOrder.Order)
				fmt.Println("Order placed successfully!")
			}
		case 2:
			fmt.Println("View available coffee drinks:")
			coffeeDrinkTypes := []string{"espresso", "latte", "cappuccino", "americano", "mocha", "macchiato"}
			for _, drinkType := range coffeeDrinkTypes {
				coffeeDrink, err := getCoffeeDrink(drinkType)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Coffee Drink Name:", coffeeDrink.getName())
					fmt.Println("Coffee Drink Strength:", coffeeDrink.getStrength())
					fmt.Println("--------")
				}
			}
		case 3:
			fmt.Println("Apply decorator to a coffee - Enter 'simple' to view simple coffee or 'cinnamon' for coffee with cinnamon:")
			var decoratorChoice string
			fmt.Scanln(&decoratorChoice)

			if decoratorChoice == "simple" {
				simpleCoffee := SimpleCoffee{Ingredients: []string{"coffee"}}
				fmt.Println("Simple Coffee Ingredients:", simpleCoffee.GetIngredients())
			} else if decoratorChoice == "cinnamon" {
				simpleCoffee := SimpleCoffee{Ingredients: []string{"coffee"}}
				cinnamonCoffee := CinnamonDecorator{BaseDecorator{&simpleCoffee}}
				fmt.Println("Coffee with Cinnamon Ingredients:", cinnamonCoffee.GetIngredients())
			} else {
				fmt.Println("Invalid choice!")
			}
		case 4:
			fmt.Println("Exiting the cafe. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
