package main

import "testing"

//Teste func
func TestSum(t *testing.T) {
	result := sum(3, 3)

	if result != 6 {
		t.Error("The result must be 6")
	}
}
