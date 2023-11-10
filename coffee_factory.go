package main

import (
	"fmt"
)

type ICoffeeDrink interface {
	setName(name string)
	setStrength(strength int)
	getName() string
	getStrength() int
}

type CoffeeDrink struct {
	name     string
	strength int
}

func (c *CoffeeDrink) setName(name string) {
	c.name = name
}

func (c *CoffeeDrink) getName() string {
	return c.name
}

func (c *CoffeeDrink) setStrength(strength int) {
	c.strength = strength
}

func (c *CoffeeDrink) getStrength() int {
	return c.strength
}

type Espresso struct {
	CoffeeDrink
}

func newEspresso() ICoffeeDrink {
	return &Espresso{
		CoffeeDrink: CoffeeDrink{
			name:     "Espresso",
			strength: 5,
		},
	}
}

type Latte struct {
	CoffeeDrink
}

func newLatte() ICoffeeDrink {
	return &Latte{
		CoffeeDrink: CoffeeDrink{
			name:     "Latte",
			strength: 3,
		},
	}
}

type Cappuccino struct {
	CoffeeDrink
}

func newCappuccino() ICoffeeDrink {
	return &Cappuccino{
		CoffeeDrink: CoffeeDrink{
			name:     "Cappuccino",
			strength: 4,
		},
	}
}

type Americano struct {
	CoffeeDrink
}

func newAmericano() ICoffeeDrink {
	return &Americano{
		CoffeeDrink: CoffeeDrink{
			name:     "Americano",
			strength: 2,
		},
	}
}

type Mocha struct {
	CoffeeDrink
}

func newMocha() ICoffeeDrink {
	return &Mocha{
		CoffeeDrink: CoffeeDrink{
			name:     "Mocha",
			strength: 3,
		},
	}
}

type Macchiato struct {
	CoffeeDrink
}

func newMacchiato() ICoffeeDrink {
	return &Macchiato{
		CoffeeDrink: CoffeeDrink{
			name:     "Macchiato",
			strength: 4,
		},
	}
}

func getCoffeeDrink(drinkType string) (ICoffeeDrink, error) {
	switch drinkType {
	case "espresso":
		return newEspresso(), nil
	case "latte":
		return newLatte(), nil
	case "cappuccino":
		return newCappuccino(), nil
	case "americano":
		return newAmericano(), nil
	case "mocha":
		return newMocha(), nil
	case "macchiato":
		return newMacchiato(), nil
	default:
		return nil, fmt.Errorf("Wrong coffee drink type passed")
	}
}
