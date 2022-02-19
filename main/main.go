package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	autoRate   = 8
	Rate       = 12
	upLimSum   = 20000000
	downLimSum = 1000000
	sumStr     = "Сумма кредита"
	timeStr    = "Срок кредита"
	percentStr = "Процентная ставка"
	firPay     = "Первоначальный взнос"
	monStr     = "Ежемесячный платеж составит"
)

func main() {
	//Запрашиваем у клиента сумму кредита
	fmt.Print("Введите сумму кредита от 1 до 20 млн:")
	sum := UserInsertData()
	if sum > upLimSum || sum < downLimSum {
		for sum > upLimSum || sum < downLimSum {
			fmt.Print("Сумма превышает лимит. Введите сумму от 1 до 20 млн:")
			sum = UserInsertData()
		}
	}
	//Запрашиваем срок кредита в годах
	fmt.Print("Введите срок кредита от 1 до 15 лет:")
	time := UserInsertData()
	if time < 1 || time > 15 {
		for time < 1 || time > 15 {
			fmt.Print("Введите корректный срок кредита:")
			time = UserInsertData()
		}
	}
	//Запрашиваем ежемесячную ставку
	fmt.Print("Введите первоначальный взнос:")
	maxFP := sum / 2
	firstPay := UserInsertData()
	for firstPay < 0 || firstPay > maxFP {
		fmt.Print("Введите корректный первоначальный взнос:")
		firstPay = UserInsertData()
	}

	//запрашиваем цель кредита
	fmt.Print("Введите цель кредита?\nЕсли на автомобиль - введите 1\nЕсли на квартиру - введите 2\n")
	goal := UserInsertData()
	if goal == 1 {
		monPay := payForMonth(sum, autoRate, time, firstPay)
		for _, s := range makeReport(sum, autoRate, time, firstPay, monPay) {
			fmt.Print(s)
		}
	} else if goal == 2 {
		monPay := payForMonth(sum, Rate, time, firstPay)
		for _, s := range makeReport(sum, Rate, time, firstPay, monPay) {
			fmt.Print(s)
		}

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
func payForMonth(sum int, rate float64, time int, firstPay int) float64 {
	convTimeInMonth := float64(time) * 12                          //конвертируем срок из лет в месяцы
	remMon := sum - firstPay                                       //Считаем остаток в соответствии с первоначальным взносом
	partF := 1 - math.Pow(1+(12/(100*rate)), convTimeInMonth*(-1)) //Рассчитываем знаменатель формулы
	monPay := float64(remMon) * (12 / (100 * rate)) / partF        //Используем полностью формулу

	return monPay
}
func makeReport(sum int, rate float64, time int, firstPay int, monPay float64) []string {
	report := []string{}
	report = append(report, fmt.Sprintf("%30s: %v\n", sumStr, sum))
	report = append(report, fmt.Sprintf("%30s: %v\n", timeStr, time))
	report = append(report, fmt.Sprintf("%30s: %v\n", percentStr, rate))
	report = append(report, fmt.Sprintf("%30s: %v\n", firPay, firstPay))
	report = append(report, "----------------------------------------------------------")
	report = append(report, fmt.Sprintf("%30s: %.2f\n", monStr, monPay))

	return report
}
