package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const upLimSum = 20000000
const downLimSum = 1000000

/*const autoRate = 0.08
const Rate = 0.12*/

func main() {
	/////////////////////////////////
	fmt.Print("Введите сумму кредита от 1 до 20 млн:")
	sum := UserInsertData()
	if sum > upLimSum || sum < downLimSum {
		for sum > upLimSum || sum < downLimSum {
			fmt.Print("Сумма превышает лимит. Введите сумму от 1 до 20 млн:")
			sum = UserInsertData()
		}
	}
	//////////////////////////
	fmt.Print("Введите срок кредита от 1 до 15 лет:")
	time := UserInsertData()
	if time < 1 || time > 15 {
		for time < 1 || time > 15 {
			fmt.Print("Введите корректный срок кредита:")
			time = UserInsertData()
		}
	}
	/////////////////////////////////////////
	fmt.Print("Введите первоначальный взнос:")
	maxFP := sum / 2
	firstPay := UserInsertData()
	for firstPay < 0 || firstPay > maxFP {
		fmt.Print("Введите корректный первоначальный взнос:")
		firstPay = UserInsertData()
	}

	//////////////////////////////////////////
	fmt.Print("Введите цель кредита?\nЕсли на автомобиль - введите 1\nЕсли на квартиру - введите 2")
	goal := UserInsertData()
	if goal == 1 {
		//////
	} else if goal == 2 {

	} else {
		fmt.Print("Введите корректный номер цели:")
		goal = UserInsertData()
	}

}
func UserInsertData() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	data, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
