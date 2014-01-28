package fizzbuzz

import (
	"testing"
)

func TestShouldGetFIZZ(t * testing.T){
	fb := NewFizzBuzz()
	if result, _ := fb.Process(3); result != FIZZ {
		t.Errorf("For num %v gets %v instead of %v", 3, result, FIZZ)
	}
}

func TestShouldGetNumber(t *testing.T){
	fb := NewFizzBuzz()
	if result, _ := fb.Process(1); result != "1" {
		t.Errorf("For num %v gets %v instead of %v", 1, result, "1")
	}
}

