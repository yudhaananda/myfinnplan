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
	Month               map[string][]entity.Transaction
	MonthTotal          map[string]float64
	MonthTotalNormalize map[string]map[string]float64
}

func FormatTransaction(transaction []entity.Transaction) TransactionFormatter {
	var result TransactionFormatter

	result.Week = make(map[string][]entity.Transaction)
	result.Month = make(map[string][]entity.Transaction)
	result.WeekTotal = make(map[string]float64)
	result.MonthTotal = make(map[string]float64)
	result.WeekTotalNormalize = make(map[string]map[string]float64)
	result.MonthTotalNormalize = make(map[string]map[string]float64)

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

	result.WeekTotalNormalize = normalizeWeek(result.MonthTotal, result.WeekTotal)
	result.MonthTotalNormalize = normalizeMonth(result.MonthTotal)

	return result
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
		temp[monthKey] = make(map[string]float64)
		for weekKey, week := range weekTotal {
			if strings.Contains(weekKey, monthKey) {
				temp[monthKey][weekKey] = week
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
