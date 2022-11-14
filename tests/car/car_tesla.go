package car

import "fmt"

type Tesla struct{}

func (m *Tesla) Run(speed int, location string) bool {
	fmt.Println("Tesla", speed, location)
	return true
}
