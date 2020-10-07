package service2

import "testing"

func TestSum(t *testing.T) {

	result := sum(1, 3)

	if result != 4 {
		t.Error("wrong!")
	}

}

//func TestPow(t *testing.T) {
//
//	result := pow(2, 4)
//
//	if result != 16 {
//		t.Errorf("wrong, got %v!", result)
//	}
//}
