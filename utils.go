package myzo

import (
	"log"
	"strconv"
)

// So this function is because Monzo stores balances, etc as a 64 bit integer.
// With that being said, this function makes it more friendly to read.
func Convert64IntToFloat(v int64) float64 {
	before := v
	if v < 0 {
		v = (v - v) - v
	}
	lastTwoDigits := v % 100
	allDigitsExceptLast := v / 100
	var lastTwoDigitsStr string
	lastTwoDigitsStr = strconv.Itoa(int(lastTwoDigits))
	if lastTwoDigits < 10 {
		lastTwoDigitsStr = "0" + lastTwoDigitsStr
	}
	allDigitsExceptLastStr := strconv.Itoa(int(allDigitsExceptLast))
	floatingPointNumber := allDigitsExceptLastStr + "." + lastTwoDigitsStr
	num, err := strconv.ParseFloat(floatingPointNumber, 64)
	if err != nil {
		log.Println("Could not convert to float!")
		return 0
	}
	if before > 0 {
		return num
	} else {
		return -num
	}
}

// So this function is because Monzo stores balances, etc as a 64 bit integer.
// With that being said, this function makes it more friendly to read.
func Convert64IntToString(v int64) string {
	before := v
	if v < 0 {
		v = (v - v) - v
	}
	lastTwoDigits := v % 100
	allDigitsExceptLast := v / 100
	var lastTwoDigitsStr string
	lastTwoDigitsStr = strconv.Itoa(int(lastTwoDigits))
	if lastTwoDigits < 10 {
		lastTwoDigitsStr = "0" + lastTwoDigitsStr
	}
	allDigitsExceptLastStr := strconv.Itoa(int(allDigitsExceptLast))
	floatingPointNumber := allDigitsExceptLastStr + "." + lastTwoDigitsStr
	if before > 0 {
		return floatingPointNumber
	} else {
		return "-" + floatingPointNumber
	}
}
