package car

import "testing"

func TestRun(t *testing.T) {
	c := NewCar()
	a := c.Run(123, "sz")
	t.Log(a)
}
