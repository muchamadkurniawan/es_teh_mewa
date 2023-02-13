package helper

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strconv"
	"time"
)

func Add(a int, b int) int {
	return a + b
}

func FormatTanggal(tanggal time.Time) string {

	t := strconv.FormatInt(int64(tanggal.Year()), 10)
	m := tanggal.Month().String()
	d := strconv.FormatInt(int64(tanggal.Day()), 10)
	return d + " " + m + " " + t
	//return tanggal.String()
}

func FormatTanggalUpdate(tanggal time.Time) string {

	t := strconv.FormatInt(int64(tanggal.Year()), 10)
	m := ""
	if tanggal.Month() > 9 {
		m = strconv.Itoa(int(tanggal.Month()))
	} else {
		m = "0" + strconv.Itoa(int(tanggal.Month()))
	}
	d := ""
	if tanggal.Day() > 9 {
		d = strconv.FormatInt(int64(tanggal.Day()), 10)
	} else {
		d = "0" + strconv.FormatInt(int64(tanggal.Day()), 10)
	}
	return t + "-" + m + "-" + d
	//return tanggal.String()
}

func FormatNumberic(number int) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprint(number)
}

func CheckBool(b bool) string {
	if b {
		return "on"
	} else {
		return ""
	}
}
