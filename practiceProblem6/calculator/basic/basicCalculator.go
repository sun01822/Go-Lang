package basic


type BasicCalculator struct {
	FirstNumber float64
	SecondNumber float64
}

func (bc BasicCalculator) Add() float64 {
	return bc.FirstNumber + bc.SecondNumber
}

func (bc BasicCalculator) Subtract() float64 {
	return bc.FirstNumber - bc.SecondNumber
}

func (bc BasicCalculator) Multiply() float64 {
	return bc.FirstNumber * bc.SecondNumber
}

func (bc BasicCalculator) Divide() float64 {
	return bc.FirstNumber / bc.SecondNumber
}

