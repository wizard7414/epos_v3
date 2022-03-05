package utils

import (
	"strconv"
	"time"
)

const IsoFormat = "2006-01-02T15:04:05-07:00"

var weekDayStr = [...]string{
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
	"Воскресенье",
}

var monthStr = [...]string{
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

func GetWeekDayStr(p time.Weekday) string {
	return weekDayStr[p]
}

func GetMonthStr(p time.Month) string {
	return monthStr[p]
}

func GetIsoDate() string {
	t := time.Now()
	return t.Format(IsoFormat)
}

func GetZimCreateDate() string {
	t := time.Now()
	weekDay := GetWeekDayStr(t.Weekday())
	day := strconv.Itoa(t.Day())
	month := GetMonthStr(t.Month())
	year := strconv.Itoa(t.Year())
	return "Создан " + weekDay + " " + day + " " + month + " " + year
}
