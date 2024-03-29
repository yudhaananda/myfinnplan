package formatter

import (
	"myfinnplan/entity"
	"strconv"
	"strings"
	"time"
)

type TransactionFormatter struct {
	Week                map[string][]entity.Transaction
	WeekTotal           map[string]float64
	WeekTotalNormalize  map[string]map[string]float64
	WeekEstimate        float64
	Month               map[string][]entity.Transaction
	MonthTotal          map[string]float64
	MonthTotalNormalize map[string]map[string]float64
	MonthEstimate       float64
}

func FormatTransaction(transaction []entity.Transaction) TransactionFormatter {
	var result TransactionFormatter

	result.Week = make(map[string][]entity.Transaction)
	result.Month = make(map[string][]entity.Transaction)
	result.WeekTotal = make(map[string]float64)
	result.MonthTotal = make(map[string]float64)
	result.WeekTotalNormalize = make(map[string]map[string]float64)
	result.MonthTotalNormalize = make(map[string]map[string]float64)
	totalAmount := 0.0
	for _, value := range transaction {
		year, week := value.CreatedDate.ISOWeek()
		if year < value.CreatedDate.Year() {
			week = 0
		}

		temp := value.CreatedDate
		for year > value.CreatedDate.Year() {
			temp = temp.AddDate(0, 0, -1)
			year, week = temp.ISOWeek()
			week += 1
		}
		yearMin, weekMin := time.Date(value.CreatedDate.Year(), value.CreatedDate.Month(), 1, 0, 0, 0, 0, time.Local).ISOWeek()
		if yearMin < value.CreatedDate.Year() {
			weekMin = 0
		}
		weekName := "week " + strconv.Itoa(week-weekMin+1) + " " + value.CreatedDate.Month().String() + " " + strconv.Itoa(value.CreatedDate.Year())
		monthName := value.CreatedDate.Month().String() + " " + strconv.Itoa(value.CreatedDate.Year())
		if weekName == "week 1 January 1" || monthName == "January 1" {
			break
		}
		result.Week[weekName] = append(result.Week[weekName], value)
		result.WeekTotal[weekName] += value.Amount

		result.Month[monthName] = append(result.Month[monthName], value)
		result.MonthTotal[monthName] += value.Amount
		totalAmount += value.Amount
	}
	monthTotal := make(map[string]float64)
	for key, value := range result.MonthTotal {
		monthTotal[key] = value
	}
	weekTotal := make(map[string]float64)
	for key, value := range result.WeekTotal {
		weekTotal[key] = value
	}
	if transaction[0].BankAccount.IsDebit {
		result.WeekEstimate, result.MonthEstimate = estimate(transaction[0].BankAccount.Amount, totalAmount, transaction[0].BankAccount.ExpiredDate)
		for key := range result.MonthTotal {
			year := strings.Split(key, " ")[1]
			monthTotal["month "+year+" estimate"] = result.MonthEstimate
			weekTotal[key+" estimate"] = result.WeekEstimate
		}
	}

	result.WeekTotalNormalize = normalizeWeek(monthTotal, weekTotal)
	result.MonthTotalNormalize = normalizeMonth(monthTotal)

	return result
}

func estimate(debitAmount, totalAmount float64, expDate time.Time) (float64, float64) {
	creditAmount := debitAmount - totalAmount

	test := time.Until(expDate)
	daysUntilExp := time.Date(1, 1, 1, int(test.Hours()), 0, 0, 0, time.Local).YearDay()

	amountPerDays := creditAmount / float64(daysUntilExp)
	amountPerWeek := creditAmount
	amountPerMonth := creditAmount
	if daysUntilExp > 7 {
		amountPerWeek = amountPerDays * 7
	}
	if daysUntilExp > 30 {
		amountPerMonth = amountPerDays * 30
	}
	return amountPerWeek, amountPerMonth
}

func normalizeMonth(monthTotal map[string]float64) (temp map[string]map[string]float64) {
	temp = make(map[string]map[string]float64)
	for key, v := range monthTotal {
		year := strings.Split(key, " ")[1]
		if temp[year] == nil {
			temp[year] = make(map[string]float64)
		}
		temp[year][key] = v
	}
	for key, val := range temp {
		_, max := findMinAndMax(val)
		for idx, v := range val {
			temp[key][idx] = ((v - 0) / (max - 0)) * 10
		}
	}
	return temp
}

func normalizeWeek(monthTotal map[string]float64, weekTotal map[string]float64) (temp map[string]map[string]float64) {
	temp = make(map[string]map[string]float64)
	for monthKey := range monthTotal {
		if !strings.Contains(monthKey, "estimate") {
			temp[monthKey] = make(map[string]float64)
			for weekKey, week := range weekTotal {
				if strings.Contains(weekKey, monthKey) {
					temp[monthKey][weekKey] = week
				}
			}
		}
	}
	for key, val := range temp {
		_, max := findMinAndMax(val)
		for idx, v := range val {
			temp[key][idx] = ((v - 0) / (max - 0)) * 10
		}
	}
	return temp
}

func findMinAndMax(a map[string]float64) (min float64, max float64) {
	i := 0
	for _, value := range a {
		if i == 0 {
			max = value
			min = value
		}
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		i++
	}
	return min, max
}
