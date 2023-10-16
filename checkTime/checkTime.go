package checkTime

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func CheckMonth() time.Month {
	month := time.Now().Month()
	return month
}
func DaysInMonth(month int) int {
	switch {
	case month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12:
		return 31
	case month == 4 || month == 6 || month == 9 || month == 11:
		return 30
	case month == 2:
		year := time.Now().Year()
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			return 29
		} else {
			return 28
		}
	}
	return 31
}

func NumericMonthes() tgbotapi.InlineKeyboardMarkup {
	var (
		numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Январь", "01"),
				tgbotapi.NewInlineKeyboardButtonData("Февраль", "02"),
				tgbotapi.NewInlineKeyboardButtonData("Март", "03"),
				tgbotapi.NewInlineKeyboardButtonData("Апрель", "04"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Май", "05"),
				tgbotapi.NewInlineKeyboardButtonData("Июнь", "06"),
				tgbotapi.NewInlineKeyboardButtonData("Июль", "07"),
				tgbotapi.NewInlineKeyboardButtonData("Август", "08"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Сентябрь", "09"),
				tgbotapi.NewInlineKeyboardButtonData("Октябрь", "10"),
				tgbotapi.NewInlineKeyboardButtonData("Ноябрь", "11"),
				tgbotapi.NewInlineKeyboardButtonData("Декабрь", "12"),
			),
		)
	)
	return numericKeyboard
}
