package main

import "testing"

//Teste func
func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}
