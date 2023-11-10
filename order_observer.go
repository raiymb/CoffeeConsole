package main

import "fmt"

type Order struct {
	Coffee Coffee
	Status string
}

type Coffee struct {
	Name string
}

type OrderObserver interface {
	UpdateOrderStatus(order Order)
}

type Customer struct {
	Name string
}

func (c Customer) UpdateOrderStatus(order Order) {
	fmt.Printf("Dear %s, your order '%s' is %s\n", c.Name, order.Coffee.Name, order.Status)
}

type CoffeeOrder struct {
	Observers []OrderObserver
	Order     Order
}

func (co *CoffeeOrder) RegisterObserver(o OrderObserver) {
	co.Observers = append(co.Observers, o)
}

func (co *CoffeeOrder) NotifyObservers() {
	for _, observer := range co.Observers {
		observer.UpdateOrderStatus(co.Order)
	}
}

func (co *CoffeeOrder) ChangeStatus(status string) {
	co.Order.Status = status
	co.NotifyObservers()
}
