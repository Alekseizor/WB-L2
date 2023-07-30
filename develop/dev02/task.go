package dev02

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	invalidStringError = fmt.Errorf("некорректная строка")
)

func UnpackString(s string) (string, error) {
	var result strings.Builder
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length; i++ {
		//сначала делаем проверку, что символ не цифра
		if runes[i] >= '0' && runes[i] <= '9' {
			return "", invalidStringError
		}
		if runes[i] == '\\' {
			i++
		}
		//проверяем, что за нашим символом следует цифра
		if i+1 < length && runes[i+1] >= '0' && runes[i+1] <= '9' {
			count, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", invalidStringError
			}
			result.WriteString(strings.Repeat(string(runes[i]), count))
			i++
		} else {
			//нет числа, значит, записываем символ один раз
			result.WriteRune(runes[i])
		}
	}

	return result.String(), nil
}

func main() {
	inputs := []string{"a4bc2d5e", "abcd", "45", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, input := range inputs {
		unpacked, err := UnpackString(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("%s => %s\n", input, unpacked)
		}
	}
}
