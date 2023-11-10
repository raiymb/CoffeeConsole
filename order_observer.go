package main

import "fmt"

type Order struct {
	Coffee Coffee
	Status string
}

type OrderObserver interface {
	UpdateOrderStatus(order Order)
}

type Customer struct {
	Name string
}

func (c Customer) UpdateOrderStatus(order Order) {
	fmt.Printf("Уважаемый %s, ваш заказ '%s' имеет статус: %s\n", c.Name, order.Coffee.Name, order.Status)
}
