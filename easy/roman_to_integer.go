package easy

type RomanNumeral int

const (
	RomanNumeralI = 1
	RomanNumeralV = 5
	RomanNumeralX = 10
	RomanNumeralL = 50
	RomanNumeralC = 100
	RomanNumeralD = 500
	RomanNumeralM = 1000
)

func romanToInt(s string) int {
	result := 0

	var prev rune

	for _, r := range s {
		switch r {
		case 'I':
			result += RomanNumeralI
		case 'V':
			result += RomanNumeralV
			if prev == 'I' {
				result -= RomanNumeralI * 2
			}
		case 'X':
			result += RomanNumeralX
			if prev == 'I' {
				result -= RomanNumeralI * 2
			}
		case 'L':
			result += RomanNumeralL
			if prev == 'X' {
				result -= RomanNumeralX * 2
			}
		case 'C':
			result += RomanNumeralC
			if prev == 'X' {
				result -= RomanNumeralX * 2
			}
		case 'D':
			result += RomanNumeralD
			if prev == 'C' {
				result -= RomanNumeralC * 2
			}
		case 'M':
			result += RomanNumeralM
			if prev == 'C' {
				result -= RomanNumeralC * 2
			}
		}

		prev = r
	}

	return result
}
