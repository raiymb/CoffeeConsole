package main

import "fmt"

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

type MilkDecorator struct {
	BaseDecorator
}

func (md MilkDecorator) GetIngredients() []string {
	ingredients := md.CoffeeIngredient.GetIngredients()
	ingredients = append(ingredients, "milk")
	return ingredients
}

type ChocolateDecorator struct {
	BaseDecorator
}

func (chocDec ChocolateDecorator) GetIngredients() []string {
	ingredients := chocDec.CoffeeIngredient.GetIngredients()
	ingredients = append(ingredients, "chocolate")
	return ingredients
}

type SugarDecorator struct {
	BaseDecorator
}

func (sd SugarDecorator) GetIngredients() []string {
	ingredients := sd.CoffeeIngredient.GetIngredients()
	ingredients = append(ingredients, "sugar")
	return ingredients
}

type SyrupDecorator struct {
	BaseDecorator
	SyrupType string
}

func (syrupDec SyrupDecorator) GetIngredients() []string {
	ingredients := syrupDec.CoffeeIngredient.GetIngredients()
	ingredients = append(ingredients, fmt.Sprintf("syrup %s", syrupDec.SyrupType))
	return ingredients
}
