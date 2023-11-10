package main

type CoffeeFactory interface {
	MakeCoffee() Coffee
}

type CoffeeDetailsFactory interface {
	GetCoffeeDetails() string
}

type EspressoFactory struct{}
type LatteFactory struct{}
type CappuccinoFactory struct{}
type MochaFactory struct{}
type AmericanoFactory struct{}

type Coffee struct {
	Name        string
	Size        string
	Ingredients []string
	Intensity   string
	Sweetness   string
}

func (f EspressoFactory) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Эспрессо",
		Size:        "Маленький",
		Ingredients: []string{"кофе"},
		Intensity:   "Средний",
		Sweetness:   "Нет",
	}
}

func (f LatteFactory) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Латте",
		Size:        "Средний",
		Ingredients: []string{"кофе", "молоко"},
		Intensity:   "Низкий",
		Sweetness:   "Средний",
	}
}

func (f CappuccinoFactory) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Капучино",
		Size:        "Большой",
		Ingredients: []string{"кофе", "молоко", "пена"},
		Intensity:   "Высокий",
		Sweetness:   "Нет",
	}
}

func (f MochaFactory) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Мокко",
		Size:        "Средний",
		Ingredients: []string{"кофе", "шоколад", "молоко"},
		Intensity:   "Средний",
		Sweetness:   "Высокий",
	}
}

func (f AmericanoFactory) MakeCoffee() Coffee {
	return Coffee{
		Name:        "Американо",
		Size:        "Большой",
		Ingredients: []string{"кофе", "вода"},
		Intensity:   "Средний",
		Sweetness:   "Нет",
	}
}

func (f EspressoFactory) GetCoffeeDetails() string {
	return "Эспрессо - маленький размер, средняя интенсивность, без сладости."
}

func (f LatteFactory) GetCoffeeDetails() string {
	return "Латте - средний размер, низкая интенсивность, средняя сладость."
}

func (f CappuccinoFactory) GetCoffeeDetails() string {
	return "Капучино - большой размер, высокая интенсивность, без сладости."
}

func (f MochaFactory) GetCoffeeDetails() string {
	return "Мокко - средний размер, средняя интенсивность, высокая сладость."
}

func (f AmericanoFactory) GetCoffeeDetails() string {
	return "Американо - большой размер, средняя интенсивность, без сладости."
}
