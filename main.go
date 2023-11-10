package main

import (
	"fmt"
	"time"
)

func main() {
	cafe := GetCafeInstance()

	cafe.Menu["эспрессо"] = EspressoFactory{}
	cafe.Menu["латте"] = LatteFactory{}
	cafe.Menu["капучино"] = CappuccinoFactory{}
	cafe.Menu["мокко"] = MochaFactory{}
	cafe.Menu["американо"] = AmericanoFactory{}

	fmt.Println("Добро пожаловать в нашу кофейню! Выберите напиток из меню:")

	for name := range cafe.Menu {
		fmt.Printf("- %s\n", name)
	}
	for {
		fmt.Println("Вы хотите стандартный кофе?(да/нет) (или 'выход' для завершения)")
		var choice string
		var orderName string
		fmt.Scanln(&choice)
		if choice == "да" {
			fmt.Print("Введите название напитка, который хотите заказать (или 'выход' для завершения): ")
			fmt.Scanln(&orderName)
			if orderName == "выход" {
				break
			}

			factory, found := cafe.Menu[orderName]
			if !found {
				fmt.Println("Извините, такого напитка нет в меню.")
				continue
			}
			order := Order{
				Coffee: factory.MakeCoffee(),
				Status: "Готов к приготовлению",
			}
			fmt.Printf("Вы заказали '%s' (%s) с интенсивностью: %s и уровнем сладости: %s. Ожидайте, мы приготовим ваш напиток!\n",
				order.Coffee.Name, order.Coffee.Size, order.Coffee.Intensity, order.Coffee.Sweetness)

			time.Sleep(5 * time.Second)
			order.Status = "Ваш заказ готов!"
			fmt.Println(order.Status)
		} else if choice == "выход" {
			return
		} else if choice == "нет" {
			fmt.Print("Введите название напитка, который хотите заказать (или 'выход' для завершения): ")
			fmt.Scanln(&orderName)
			if orderName == "выход" {
				break
			}

			factory, found := cafe.Menu[orderName]
			if !found {
				fmt.Println("Извините, такого напитка нет в меню.")
				continue
			}

			order := Order{
				Coffee: factory.MakeCoffee(),
				Status: "Готов к приготовлению",
			}

			fmt.Println("Выберите размер кофе:")
			fmt.Println("1. Маленький")
			fmt.Println("2. Средний")
			fmt.Println("3. Большой")

			var orderSize int
			fmt.Print("Введите номер размера: ")
			fmt.Scanln(&orderSize)

			size := ""
			switch orderSize {
			case 1:
				size = "Маленький"
			case 2:
				size = "Средний"
			case 3:
				size = "Большой"
			default:
				fmt.Println("Неверный номер размера.")
				return
			}

			order = Order{
				Coffee: factory.MakeCoffee(),
				Status: "Готов к приготовлению",
			}

			order.Coffee.Size = size

			// Coffee Intensity Selection
			fmt.Println("Выберите интенсивность кофе:")
			fmt.Println("1. Низкая")
			fmt.Println("2. Средняя")
			fmt.Println("3. Высокая")

			var orderIntensity int
			fmt.Print("Введите номер интенсивности: ")
			fmt.Scanln(&orderIntensity)

			intensity := ""
			switch orderIntensity {
			case 1:
				intensity = "Низкий"
			case 2:
				intensity = "Средний"
			case 3:
				intensity = "Высокий"
			default:
				fmt.Println("Неверный номер интенсивности.")
				return
			}

			order.Coffee.Intensity = intensity

			// Sweetness Level Selection
			fmt.Println("Выберите уровень сладости:")
			fmt.Println("1. Нет")
			fmt.Println("2. Низкий")
			fmt.Println("3. Средний")
			fmt.Println("4. Высокий")

			var orderSweetness int
			fmt.Print("Введите номер уровня сладости: ")
			fmt.Scanln(&orderSweetness)
			order.Coffee.Sweetness = ""
			switch orderSweetness {
			case 1:
				order.Coffee.Sweetness = "Нет"
			case 2:
				order.Coffee.Sweetness = "Низкий"
			case 3:
				order.Coffee.Sweetness = "Средний"
			case 4:
				order.Coffee.Sweetness = "Высокий"
			default:
				fmt.Println("Неверный номер уровня сладости.")
				return
			}

			fmt.Print("Хотите добавить корицца? (да/нет): ")
			var addSugar string
			_, err := fmt.Scanln(&addSugar)
			if err != nil {
				return
			}
			if addSugar == "да" {
				cinnamonDecorator := CinnamonDecorator{}
				order.Coffee = cinnamonDecorator.DecorateCoffee(order.Coffee)
			}

			fmt.Printf("Вы заказали '%s' (%s) с интенсивностью: %s и уровнем сладости: %s. Ожидайте, мы приготовим ваш напиток!\nИнгридиенты:%s",
				order.Coffee.Name, order.Coffee.Size, order.Coffee.Intensity, order.Coffee.Sweetness, order.Coffee.Ingredients)

			time.Sleep(5 * time.Second)
			order.Status = "Ваш заказ готов!"
			fmt.Println(order.Status)
		} else {
			fmt.Println("Выберите один из двух вариантов!")
		}
	}
}
