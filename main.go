package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Добро пожаловать в калькулятор \nСейчас он имеет только 4 функции: +, -, *, /")
	romanToNum := RomanToNumber()

	fmt.Print("Введите первое число (римское): ")
	reader := bufio.NewReader(os.Stdin)
	firstRoman, _ := reader.ReadString('\n')
	firstRoman = strings.TrimSpace(firstRoman)

	first, err := RomanToInt(firstRoman, romanToNum)
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа:", err)
		return
	}

	fmt.Print("Введите второе число (римское): ")
	secondRoman, _ := reader.ReadString('\n')
	secondRoman = strings.TrimSpace(secondRoman)

	second, err := RomanToInt(secondRoman, romanToNum)
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа:", err)
		return
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	result, err := calculator(first, second, operator)
	if err != nil {
		fmt.Println("Ошибка при выполнении операции:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func calculator(first int, second int, sign string) (int, error) {
	switch {
	case sign == "+":
		return first + second, nil
	case sign == "-":
		return first - second, nil
	case sign == "*":
		return first * second, nil
	case sign == "/":
		if second != 0 {
			return first / second, nil
		}
		return 0, errors.New("Деление на ноль")
	default:
		return 0, errors.New("Неверный оператор")
	}
}

func RomanToNumber() map[string]int {
	RomanToNum := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return RomanToNum
}

func RomanToInt(input string, romanToNum map[string]int) (int, error) {
	var roman string
	if val, ok := romanToNum[input]; ok {
		return val, nil
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Некорректное число")
	}

	return number, nil

	if val, ok := romanToNum[roman]; ok {
		return val, nil
	}
	return 0, errors.New("Некорректное число")
}
