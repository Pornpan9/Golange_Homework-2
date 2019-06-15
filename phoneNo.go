package main

import (
	"fmt"
	"strings"
)

func phoneNormalizer(phones []string) (resultValue map[string]int) {
	resultValue = map[string]int{}
	for _, result := range phones {
		result = strings.ReplaceAll(result, " ", "")
		result = strings.ReplaceAll(result, "-", "")
		result = strings.ReplaceAll(result, "(", "")
		result = strings.ReplaceAll(result, ")", "")
		resultValue[result] = resultValue[result] + 1
	}

	return resultValue
}

func main() {
	list := []string{
		"1234567890",
		"123 456 7891",
		"123-456-7890",
		"1234567892",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"(123)456-7892",
	}
	fmt.Println(phoneNormalizer(list))

}
