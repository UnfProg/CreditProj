package main

import (
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

func main() {
	//Запрашиваем у клиента сумму кредита
	sum := insSum()
	//Запрашиваем срок кредита в годах
	time := insTime()
	//Запрашиваем первоначальный взнос
	firstPay := insertFPay(sum)
	//запрашиваем цель кредита
	rate := insRate()
	//Считаем ежемесячный платеж и ставку
	monPay := payForMonth(sum, rate, time, firstPay)
	//Формируем и выводим отчет
	for _, s := range makeReport(sum, rate, time, firstPay, monPay) {
		fmt.Print(s)
	}

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
	report = append(report, "----------------------------------------------------------\n")
	report = append(report, fmt.Sprintf("%30s: %.2f\n", monStr, monPay))
	return report
}
