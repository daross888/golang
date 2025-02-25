package fifteenth

import "strings"

type RomanNumeral struct {
	Number uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Number {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Number
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Number
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
