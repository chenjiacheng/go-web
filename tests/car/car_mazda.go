package car

import "fmt"

type Mazda struct{}

func (m *Mazda) Run(speed int, location string) bool {
	fmt.Println("Mazda", speed, location)
	return true
}
