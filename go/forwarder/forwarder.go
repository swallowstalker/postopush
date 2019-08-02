package forwarder

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// ErrIncompleteEnv error shows when one of bot params is incomplete
var ErrIncompleteEnv = errors.New("incomplete env")

// Forward forward the message to telegram bot using token and chatID
func Forward(token, message string, chatID int64) error {

	if token == "" || message == "" {
		return ErrIncompleteEnv
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = tgbotapi.ModeHTML
	_, err = bot.Send(msg)

	return err
}
