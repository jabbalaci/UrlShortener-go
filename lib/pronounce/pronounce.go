package pronounce

import (
	"fmt"
	"strings"
	"unicode"
)

var digits = map[string]string{
	"0": "zero",
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

var names = map[string]string{
	"a": "Alice",
	"b": "Benjamin",
	"c": "Charlotte",
	"d": "Daniel",
	"e": "Emily",
	"f": "Frederick",
	"g": "Grace",
	"h": "Henry",
	"i": "Isabella",
	"j": "James",
	"k": "Katherine",
	"l": "Leslie",
	"m": "Mia",
	"n": "Noah",
	"o": "Olivia",
	"p": "Penelope",
	"q": "Quinn",
	"r": "Robert",
	"s": "Sophia",
	"t": "Thomas",
	"u": "Ursula",
	"v": "Victoria",
	"w": "William",
	"x": "Xavier",
	"y": "Yara",
	"z": "Zoe",
}

func Say(text string) []string {
	result := make([]string, 0, len(text))

	for _, r := range text {
		c := string(r)
		lowercased := strings.ToLower(c)
		if digit, ok := digits[c]; ok {
			result = append(result, digit)
		} else if name, ok := names[lowercased]; ok {
			if unicode.IsLower(r) {
				name = strings.ToLower(name)
				result = append(result, fmt.Sprintf("lowercase %s", name))
			} else {
				result = append(result, fmt.Sprintf("UPPERCASE %s", name))
			}
		} else {
			result = append(result, c) // fallback case
		}
	}

	return result
}
