package tests

import "testing"

func TestGetInstance(t *testing.T) {
	s := GetInstance()
	t.Log(s)
}
