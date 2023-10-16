package monthGeneration

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram_bot/checkTime"
	api "telegram_bot/handleApi"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

type Person struct { //структура для хранения данных о клиенте
	month string
	day   int
}

func MonthCalendar(query *tgbotapi.CallbackQuery) {
	var (
		text         string
		date         string
		queryInfo    string
		daysInMonth  int
		weeksInMonth int
	)

	dayChose := tgbotapi.NewInlineKeyboardMarkup()
	m, _ := strconv.ParseInt(query.Data, 10, 8)
	daysInMonth = checkTime.DaysInMonth(int(m))
	if daysInMonth < 29 {
		weeksInMonth = 4
	} else {
		weeksInMonth = 5
	}

	for i := 1; i < weeksInMonth+1; i++ { // week
		var row []tgbotapi.InlineKeyboardButton
		for j := 1; j < 8; j++ { // day in week
			if (i-1)*7+j > daysInMonth {
				date = "\xE2\x9D\x8C"
			} else {
				date = strconv.Itoa((i-1)*7 + j)
				if len(date) <= 1 {
					queryInfo = query.Data + "0" + date
				} else {
					queryInfo = query.Data + date
				}
			}
			btn := tgbotapi.NewInlineKeyboardButtonData(date, queryInfo)
			row = append(row, btn)

		}
		dayChose.InlineKeyboard = append(dayChose.InlineKeyboard, row)
	}

	text = "Выберите день из доступных"
	msg := tgbotapi.NewMessage(query.From.ID, text)
	msg.ReplyMarkup = dayChose
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}
