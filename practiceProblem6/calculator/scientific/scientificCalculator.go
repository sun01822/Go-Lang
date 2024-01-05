package scientific


// ScientificCalculator is a struct that implements the CalculatorInterface
type ScientificCalculator struct {
	FirstNumber float64
	SecondNumber float64
}

// Add is a method that adds two numbers
func (s ScientificCalculator) Add() float64 {
	return s.FirstNumber + s.SecondNumber
}

// Subtract is a method that subtracts two numbers
func (s ScientificCalculator) Subtract() float64 {
	return s.FirstNumber - s.SecondNumber
}

// Multiply is a method that multiplies two numbers
func (s ScientificCalculator) Multiply() float64 {
	return s.FirstNumber * s.SecondNumber
}

// Divide is a method that divides two numbers
func (s ScientificCalculator) Divide() float64 {
	return s.FirstNumber / s.SecondNumber
}

