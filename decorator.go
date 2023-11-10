package main

type CinnamonDecorator struct{}

func (md CinnamonDecorator) DecorateCoffee(coffee Coffee) Coffee {
	coffee.Ingredients = append(coffee.Ingredients, "корицца")
	return coffee
}
