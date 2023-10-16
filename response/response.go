package response

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"telegram_bot/checkTime"
	"telegram_bot/dayGeneration"
	api "telegram_bot/handleApi"
	monthCalendarGeneration "telegram_bot/monthGeneration"
)

var (
	bot, _ = tgbotapi.NewBotAPI(api.GetApiToken())
)

type Person struct { //структура для хранения данных о клиенте
	month string
	day   int
}

func HandleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text
	command := message.Command()
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	var err error

	if user == nil {
		return
	}
	// Print to console
	log.Printf("%s wrote %s", user.FirstName, text)

	if command != "" {
		err = handleCommand(message.Chat.ID, command)
	} else {
		msg.Text = "Для записи используйте команды"
		_, err = bot.Send(msg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

func handleCommand(chatId int64, command string) error {
	var err error
	msg := tgbotapi.NewMessage(chatId, "")

	switch command {
	case "start":
		msg.Text = "Для записи нажмите /book"

	case "test":
		msg.Text = "not valid command 404 lol"

	case "book":
		msg.Text = "Выберете месяц"
		msg.ReplyMarkup = checkTime.NumericMonthes()

	default:
		msg.Text = "Нет такой команды"
	}
	_, err = bot.Send(msg)
	return err
}

func HandleButton(query *tgbotapi.CallbackQuery) {
	dataLen := len(query.Data)
	switch {
	// Handle month button
	case dataLen == 2:

		m, _ := strconv.ParseInt(query.Data, 10, 8) //check is month available for booking if not return
		if m < int64(checkTime.CheckMonth()) {
			msg := tgbotapi.NewMessage(query.From.ID, "Данный месяц не доступен для записи")
			if _, err := bot.Send(msg); err != nil {
				panic(err) // not correct way handle error, remake!
			}
			return
		}
		monthCalendarGeneration.MonthCalendar(query)
		break

	// Handle day button
	case dataLen > 2 && dataLen < 5:
		dayGeneration.DayTime(query)
		break

	//handle time of the day
	case dataLen > 5:
		dayGeneration.ApproveBooking(query)
		break
	default:
		msg := tgbotapi.NewMessage(query.From.ID, "Неопознаная кнопка")
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
		break
	}

}
