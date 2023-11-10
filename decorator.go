package main

type CoffeeIngredient interface {
	GetIngredients() []string
}

type SimpleCoffee struct {
	Ingredients []string
}

func (sc SimpleCoffee) GetIngredients() []string {
	return sc.Ingredients
}

type BaseDecorator struct {
	CoffeeIngredient
}

type CinnamonDecorator struct {
	BaseDecorator
}

func (cd CinnamonDecorator) GetIngredients() []string {
	ingredients := cd.CoffeeIngredient.GetIngredients()
	ingredients = append(ingredients, "cinnamon")
	return ingredients
}
