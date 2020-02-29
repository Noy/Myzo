package monzo

import (
	"testing"
)

func TestConvertFloat(t *testing.T) {
	var negativeNumber int64 = -45569
	var positiveNumber int64 = 45569
	expectedNegative := -455.69
	expectedPositive := 455.69
	convertedNegative := Convert64IntToFloat(negativeNumber)
	convertedPositive := Convert64IntToFloat(positiveNumber)
	if convertedNegative == expectedNegative && convertedPositive == expectedPositive {
		t.Logf("It worked. %v was positive & %v was negative", convertedPositive, convertedNegative)
	} else {
		t.Error("Did not work.. check the func: ", convertedPositive, convertedNegative)
	}
}