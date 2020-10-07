package service1

import "testing"

func TestSum(t *testing.T) {

	result := sum(1, 3)

	if result != 4 {
		t.Error("wrong!")
	}

}

func TestMultiply(t *testing.T) {

	result := multiply(1, 3)

	if result != 3 {
		t.Error("wrong!")
	}

}
