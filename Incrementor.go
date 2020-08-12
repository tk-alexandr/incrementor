package incrementor

import "math"

//Incrementor type. Instances should be created with "CreateIncrementor" function
type Incrementor struct {
	currentNumber int
	maxValue int 
}

//CreateIncrementor function wich creates instances of type "Incrementor"
func CreateIncrementor () *Incrementor {
	return &Incrementor { maxValue: math.MaxInt64 }
}

//GetNumber returns current value of Incrementor's instance
func (i *Incrementor) GetNumber () int {
	return i.currentNumber
}

//IncrementNumber increments current number of Incrementor's instance. If achieves max number, current number becomes 0
func (i *Incrementor) IncrementNumber () {	
	i.currentNumber++

	if i.currentNumber >= i.maxValue {
		i.currentNumber = 0
	}
}

//SetMaxValue sets max value. If current number overflows max value, current number becomes 0
func (i *Incrementor) SetMaxValue (maxValue int) {

	if maxValue <= 0 {
		panic("Max Value of Incrementor can't be above 0")
	}

	if maxValue >= i.currentNumber {
		i.currentNumber = 0
	}

	i.maxValue = maxValue
}