package car

import "sync"

type Car struct {
	Interface Interface
}

var once sync.Once

var internalCar *Car

func NewCar() *Car {
	once.Do(func() {
		internalCar = &Car{Interface: &Tesla{}}
	})

	return internalCar
}

func (c *Car) Run(speed int, location string) bool {
	return c.Interface.Run(speed, location)
}
