package main

import "sync"

type Cafe struct {
	Menu      map[string]ICoffeeDrink
	Orders    []Order
	Observers []OrderObserver
}

var cafeInstance *Cafe
var once sync.Once

func GetCafeInstance() *Cafe {
	once.Do(func() {
		cafeInstance = &Cafe{
			Menu: make(map[string]ICoffeeDrink),
		}
	})
	return cafeInstance
}
