package fizzbuzz

import (
	"strconv"
)

const (
	FIZZ = "FIZZ"
	BUZZ = "BUZZ"
	FIZZ_NUMBER = 3
)

type FizzBuzz struct {
}

func (self *FizzBuzz) IsFizz(number int) bool {
	return (number % FIZZ_NUMBER == 0)
}

func (self *FizzBuzz) Process(number int) (string, error) {
	if ok :=  self.IsFizz(number); ok {
		return FIZZ, nil
	}
		
	return strconv.Itoa(number), nil
}

func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{}
}
