package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInsertData() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	data, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return data, nil
}
func insTime() int {
	fmt.Print("Введите срок кредита от 1 до 15 лет:")
	time, err := UserInsertData()
	for time < 1 || time > 15 || err != nil {
		fmt.Print("Введите корректный срок кредита:")
		time, err = UserInsertData()
	}

	return time
}
func insertFPay(sum int) int {
	fmt.Print("Введите первоначальный взнос:")
	maxFP := sum / 2
	firstPay, err := UserInsertData()
	for firstPay < 0 || firstPay > maxFP || err != nil {
		fmt.Print("Введите корректный первоначальный взнос:")
		firstPay, err = UserInsertData()

	}
	return firstPay
}
func insRate() float64 {
	fmt.Print("Введите цель кредита?\nЕсли на автомобиль - введите 1\nЕсли на квартиру - введите 2\n")
	goal, err := UserInsertData()
	for err != nil || (goal < 1 || goal > 2) {
		fmt.Print("Введите корректный номер цели:")
		goal, err = UserInsertData()
	}
	if goal == 1 {
		return autoRate
	} else if goal == 2 {
		return Rate
	}
	return 0
}
func insSum() int {
	fmt.Print("Введите сумму кредита от 1 до 20 млн:")
	sum, err := UserInsertData()
	for sum > upLimSum || sum < downLimSum || err != nil {
		fmt.Print("Сумма превышает лимит. Введите сумму от 1 до 20 млн:")
		sum, err = UserInsertData()
	}
	return sum
}
