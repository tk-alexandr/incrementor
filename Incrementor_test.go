package incrementor

import (
	"testing"
	"math"
)

func TestGetNumberReturnsZeroWhenNotIncremented (t *testing.T) {
	i := CreateIncrementor()
	
	val := i.GetNumber()
	if val != 0 {
		t.Error("Not incremented Incrementor should return 0, got ", val)
	}
}


func TestIncrementNumber (t *testing.T) {
	i := CreateIncrementor()
	
	i.IncrementNumber()
	i.IncrementNumber()

	val := i.GetNumber()
	if val != 2 {
		t.Error("currentNumber should turn 2 when incremented 2 times, but got ", val)
	}
}

func TestIncrementNumberIfMaxValueNotSetShouldSetMaxValue (t *testing.T) {
	i := CreateIncrementor()
	

	val := i.maxValue
	if val != math.MaxInt64 {
		t.Error("Max value should be MaxInt64 if not set, got ", val)
	}
}

func TestSetMaxValueCurrentNumberShouldTurnZeroWhenMaxValueAchieved (t *testing.T) {
	i := CreateIncrementor()
	
	exp1 := 20
	exp2 := 0
	i.SetMaxValue(21)
		
	for j := 0; j < exp1; j++ {
		i.IncrementNumber()
	}
	
	val1 := i.GetNumber()
	i.IncrementNumber()
	val2 := i.GetNumber()

	if !(val1 == exp1 && val2 == exp2) {
		t.Errorf("When achieving max value: expected pre-max value is %d, got %d. When max value achieved expected %d, got %d", exp1, val1, exp2, val2)
	}
}