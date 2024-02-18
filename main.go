package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var SelectOperand = map[string]string{
	"+": "+",
	"-": "-",
	"/": "/",
	"*": "*",
}

// Преобразование реультата в римскую систему исчисления
func convertToRoman(romanResult int) (string, error) {
	var romanNum string
	if romanResult == 0 {
		return "", errors.New("it can't be zero")
	} else if romanResult < 0 {
		return "", errors.New("it can't be negative")
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	return romanNum, nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	result, mode, err := Result(parts)
	if err != nil {
		panic(err)
	}
	// Проверяем какая система была в запросе и выдаем соответствующий ответ
	switch mode {
	case 1:
		text, err := convertToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	case 2:
		fmt.Println(result)
	}

}

func Result(parts []string) (int, int, error) {
	// mode передает в main какаяистема исчисления используется
	var mode int
	// Проверка наличия 3-х параметров
	if len(parts) == 3 {
		var a, operand, b = parts[0], parts[1], parts[2]
		// Римская система
		// Проверка условий задачи от 1 до 10 включительно
		if i, err := roman[a]; err && i <= 10 && 1 <= i {
			if j, err := roman[b]; err && j <= 10 && 1 <= j {
				mode = 1
				return solution(i, j, mode, operand)
			}
		}
		// Арабская система
		// Проверка условий задачи от 1 до 10 включительно
		if i, err := strconv.Atoi(a); err == nil && i <= 10 && 1 <= i {
			if j, err := strconv.Atoi(b); err == nil && j <= 10 && 1 <= j {
				mode = 2
				return solution(i, j, mode, operand)
			} else {
				return 0, mode, errors.New("the second parameter is not correct")
			}
		} else {
			return 0, mode, errors.New("the first parameter is not correct")
		}
	} else {
		return 0, mode, errors.New("invalid operation")
	}

}

// Арифметическая функция
func solution(a, b, mode int, o string) (int, int, error) {
	switch o {
	case "+":
		return a + b, mode, nil
	case "-":
		return a - b, mode, nil
	case "/":
		return a / b, mode, nil
	case "*":
		return a * b, mode, nil
	}
	return 0, mode, errors.New("operand is invalid")
}
