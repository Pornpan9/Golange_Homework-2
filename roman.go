package main

import "fmt"

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value

			// fmt.Println("Print Value Test : ", number, roman)
		}
	}
	return roman
}

func main() {

	for i := 1; i <= 100; i++ {
		// fmt.Println("Print i = ", i)
		r := Roman(i)
		fmt.Println(i, "==", r)
	}

}
