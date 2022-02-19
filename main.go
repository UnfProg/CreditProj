package main

import (
	"awesomeProject/messenger"
	"fmt"
	"math"
)

const (
	autoRate   = 8.0
	Rate       = 12.0
	upLimSum   = 20000000
	downLimSum = 1000000
	sumStr     = "Сумма кредита"
	timeStr    = "Срок кредита"
	percentStr = "Процентная ставка"
	firPay     = "Первоначальный взнос"
	monStr     = "Ежемесячный платеж составит"
)

var myIO = messenger.MessengerConsole{}

//var myIO = messenger.MessengerTelegram{}

func main() {

	//Запрашиваем у клиента сумму кредита
	sum := AskSum()
	//Запрашиваем срок кредита в годах
	time := AskTime()
	//Запрашиваем первоначальный взнос
	firstPay := AskFirstPay(sum)
	//запрашиваем цель кредита
	rate := AskPurpose()
	//Считаем ежемесячный платеж и ставку
	monPay := payForMonth(sum, rate, time, firstPay)
	//Формируем и выводим отчет
	report := makeReport(sum, rate, time, firstPay, monPay)
	myIO.SendMultiline(report)
}

func AskSum() int {
	myIO.SendMessage("Введите сумму кредита от 1 млн до 20 млн:")
	sum, err := myIO.ReadData()
	for sum > upLimSum || sum < downLimSum || err != nil {
		myIO.SendMessage("Сумма превышает лимит. Введите сумму от 1 до 20 млн:")
		sum, err = myIO.ReadData()
	}
	return sum
}

func AskTime() int {
	myIO.SendMessage("Введите срок кредита от 1 до 15 лет:")
	time, err := myIO.ReadData()
	for time < 1 || time > 15 || err != nil {
		myIO.SendMessage("Введите корректный срок кредита:")
		time, err = myIO.ReadData()
	}

	return time
}

func AskFirstPay(sum int) int {
	myIO.SendMessage("Введите первоначальный взнос:")
	maxFP := sum / 2
	firstPay, err := myIO.ReadData()
	for firstPay < 0 || firstPay > maxFP || err != nil {
		myIO.SendMessage("Введите корректный первоначальный взнос:")
		firstPay, err = myIO.ReadData()

	}
	return firstPay
}

func AskPurpose() float64 {
	myIO.SendMessage("Введите цель кредита?\nЕсли на автомобиль - введите 1\nЕсли на квартиру - введите 2")
	goal, err := myIO.ReadData()
	for err != nil || (goal < 1 || goal > 2) {
		myIO.SendMessage("Введите корректный номер цели:")
		goal, err = myIO.ReadData()
	}
	if goal == 1 {
		return autoRate
	} else if goal == 2 {
		return Rate
	}
	return 0
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
	report = append(report, fmt.Sprintf("%30s: %v рублей", sumStr, sum))
	report = append(report, fmt.Sprintf("%30s: %v год(а)", timeStr, time))
	report = append(report, fmt.Sprintf("%30s: %v процентов", percentStr, rate))
	report = append(report, fmt.Sprintf("%30s: %v рублей", firPay, firstPay))
	report = append(report, "----------------------------------------------------------")
	report = append(report, fmt.Sprintf("%30s: %.2f рублей", monStr, monPay))
	return report
}
