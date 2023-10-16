package dayGeneration

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	api "telegram_bot/handleApi"
	"time"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

func DayTime(query *tgbotapi.CallbackQuery) {

	workHours := 10.0

	timePerClient := 1.5
	timeBetweenClient := 0.5

	dataInt, _ := strconv.ParseInt(query.Data, 10, 32)
	date := dataInt % 100
	month := dataInt / 100
	text := "Расписание на " + strconv.Itoa(int(date)) + " число " + strconv.Itoa(int(month)) + " месяца"
	startHour := time.Date(2023, time.Month(month), int(date), 10, 00, 0, 0, nil)
	schedule := tgbotapi.NewInlineKeyboardMarkup()

	numPLaces := (workHours + timeBetweenClient) / (timePerClient + timeBetweenClient)

	for i := 1; i <= int(numPLaces)/3+1; i++ {
		var row []tgbotapi.InlineKeyboardButton
		for j := 1; j < 4; j++ {
			data := startHour + float64(j+(i-1)*2-2)*(timePerClient+timeBetweenClient)
			//startData := startHour + 1
			btnData := strconv.FormatFloat(data, 'f', -1, 32)
			btn := tgbotapi.NewInlineKeyboardButtonData(btnData, query.Data)
			row = append(row, btn)
		}
		schedule.InlineKeyboard = append(schedule.InlineKeyboard, row)
	}

	msg := tgbotapi.NewMessage(query.From.ID, text)
	msg.ReplyMarkup = schedule
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}

func ApproveBooking(query *tgbotapi.CallbackQuery) {
	fmt.Println(query.Data)
	text := "lol"

	msg := tgbotapi.NewMessage(query.From.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err) // not correct way handle error, remake!
	}
}
