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

func TestShouldGetBUZZ(t *testing.T){
	fb := NewFizzBuzz()
	if result, _ := fb.Process(5); result != BUZZ {
		t.Errorf("For num %v gets %v instead of %v", 5, result, BUZZ)
	}
}

func TestShouldGetFIZZBUZZ(t *testing.T){
	fb := NewFizzBuzz()
	if result, _ := fb.Process(15); result != "FizzBuzz" {
		t.Errorf("For num %v gets %v instead of %v", 15, result, "FizzBuzz")
	}
}
