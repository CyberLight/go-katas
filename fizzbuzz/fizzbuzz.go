package fizzbuzz

import (
	"strconv"
)

const (
	FIZZ = "Fizz"
	BUZZ = "Buzz"
	FIZZ_BUZZ = "FizzBuzz"
	FIZZ_NUMBER = 3
	BUZZ_NUMBER = 5
	FIZZ_BUZZ_NUMBER = FIZZ_NUMBER * BUZZ_NUMBER
)

type FizzBuzz struct {
}

func (self *FizzBuzz) IsFizz(number int) bool {
	return (number % FIZZ_NUMBER == 0)
}

func (self *FizzBuzz) IsBuzz(number int) bool {
	return (number % BUZZ_NUMBER == 0)
}

func (self *FizzBuzz) IsFizzBuzz(number int) bool {
	return (number % FIZZ_BUZZ_NUMBER == 0)
}

func (self *FizzBuzz) Process(number int) (string, error) {
	
	if ok :=  self.IsFizzBuzz(number); ok {
		return FIZZ_BUZZ, nil
	}

	if ok :=  self.IsFizz(number); ok {
		return FIZZ, nil
	}

	if ok := self.IsBuzz(number); ok {
		return BUZZ, nil
	}
		
	return strconv.Itoa(number), nil
}

func NewFizzBuzz() *FizzBuzz {
	return &FizzBuzz{}
}
