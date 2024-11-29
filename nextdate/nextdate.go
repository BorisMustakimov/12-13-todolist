package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {

	var nextDate time.Time

	parsedDate, err := time.Parse(DateFormat, date)
	if err != nil {
		return "", fmt.Errorf("неверный формат даты")
	}

	if repeat == "" {
		return "", fmt.Errorf("неверный формат даты")
	}

	dateRepeat := strings.Split(repeat, " ")

	switch dateRepeat[0] {

	case "d":
		days, err := strconv.Atoi(dateRepeat[1])
		if err != nil || days < 1 || days > 400 {
			return "", fmt.Errorf("указано неверное число (необходимо 1-400)")
		}
		if len(dateRepeat) == 0 {
			return "", fmt.Errorf("не указано повторение")
		}
		if parsedDate.Equal(now) {
			nextDate = now
		} else {
			nextDate = parsedDate.AddDate(0, 0, days)
			for nextDate.Before(now) {
				nextDate = nextDate.AddDate(0, 0, days)
			}
		}
	case "y":
		years, err := strconv.Atoi(dateRepeat[1])
		if err != nil {
			return "", fmt.Errorf("неверный формат даты")
		}
		if parsedDate.Equal(now) {
			nextDate = now
		} else {
			nextDate = parsedDate.AddDate(years, 0, 0)
			for nextDate.Before(now) {
				nextDate = nextDate.AddDate(years, 0, 0)
			}

		}
	case "m":
		return "", fmt.Errorf("данная функци пока не доступна")
	case "w":
		return "", fmt.Errorf("данная функци пока не доступна")
	}
	return "", nil
}
