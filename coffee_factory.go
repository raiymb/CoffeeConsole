package main

type CoffeeFactory interface {
	MakeCoffee() Coffee
}

type Espresso struct{}
type Latte struct{}
type Cappuccino struct{}
type Mocha struct{}
type Americano struct{}

type Coffee struct {
	Name        string
	Size        string
	Ingredients []string
	Intensity   string
	Sweetness   string
}

func (f Espresso) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Эспрессо",
		Size:        "Маленький",
		Ingredients: []string{"кофе"},
		Intensity:   "Средний",
		Sweetness:   "Нет",
	}
}

func (f Latte) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Латте",
		Size:        "Средний",
		Ingredients: []string{"кофе", "молоко"},
		Intensity:   "Низкий",
		Sweetness:   "Средний",
	}
}

func (f Cappuccino) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Капучино",
		Size:        "Большой",
		Ingredients: []string{"кофе", "молоко", "пена"},
		Intensity:   "Высокий",
		Sweetness:   "Нет",
	}
}

func (f Mocha) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Мокко",
		Size:        "Средний",
		Ingredients: []string{"кофе", "шоколад", "молоко"},
		Intensity:   "Средний",
		Sweetness:   "Высокий",
	}
}

func (f Americano) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Американо",
		Size:        "Большой",
		Ingredients: []string{"кофе", "вода"},
		Intensity:   "Средний",
		Sweetness:   "Нет",
	}
}
