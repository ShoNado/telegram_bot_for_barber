package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Январь", "депс"),
		tgbotapi.NewInlineKeyboardButtonData("Февраль", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Март", "3"),
		tgbotapi.NewInlineKeyboardButtonData("Апрель", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Май", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Июнь", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Июль", "6"),
		tgbotapi.NewInlineKeyboardButtonData("Август", "8"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сентябрь", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Октябрь", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Ноябрь", "6"),
		tgbotapi.NewInlineKeyboardButtonData("Декабрь", "3"),
	),
)

func main() {
	TelegramApitoken := "6371492060:AAGJvIz6K30ooHiQc5L61lU1mOp9vuhkOsY"
	bot, err := tgbotapi.NewBotAPI(TelegramApitoken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		caseAnswer()
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			// If the message was open, add a copy of our numeric keyboard.
			switch update.Message.Text {
			case "/start":
				msg.Text = "Для записи нажмите /book"
			case "/test":
				msg.ReplyMarkup = numericKeyboard
				msg.Text = "Доступные варианты"
			case "/book":
				_, month, day := time.Now().Date()

				fmt.Printf("month = %v\n", month)
				fmt.Printf("day = %v\n", day)
				msg.ReplyMarkup = numericKeyboard
				msg.Text = "Выберете месяц"
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
