package bot

import (
	"github.com/getsentry/sentry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("reported to Sentry: %s", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var msg tgbotapi.MessageConfig
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg, err = HandleState(update, msg)
		if err != nil {
			sentry.CaptureException(err)
			log.Printf("reported to Sentry: %s", err)
			return
		}

		if _, err = bot.Send(msg); err != nil {
			sentry.CaptureException(err)
			log.Printf("reported to Sentry: %s", err)
			return
		}
	}
}
