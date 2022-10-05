package formatter

import (
	"myfinnplan/entity"
	"strconv"
	"time"
)

type TransactionFormatter struct {
	Week       map[string][]entity.Transaction
	WeekTotal  map[string]float64
	Month      map[string][]entity.Transaction
	MonthTotal map[string]float64
}

func FormatTransaction(transaction []entity.Transaction) TransactionFormatter {
	var result TransactionFormatter

	result.Week = make(map[string][]entity.Transaction)
	result.Month = make(map[string][]entity.Transaction)
	result.WeekTotal = make(map[string]float64)
	result.MonthTotal = make(map[string]float64)

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

		result.Week[weekName] = append(result.Week[weekName], value)
		result.WeekTotal[weekName] += value.Amount

		result.Month[monthName] = append(result.Month[monthName], value)
		result.MonthTotal[monthName] += value.Amount

	}

	return result
}
