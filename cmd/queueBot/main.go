package main

import (
	"github.com/getsentry/sentry-go"
	_ "github.com/richard-on/QueueBot/pkg/queueBot"
	"github.com/richard-on/QueueBot/pkg/queueBot/bot"
	"github.com/richard-on/QueueBot/pkg/queueBot/db"
	"log"
	"os"
	"time"
)

var SentryDsn = os.Getenv("SENTRY_DSN")

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              SentryDsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	log.Println("Initializing Database")
	err = db.InitDb()
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("reported to Sentry: %s", err)
		return
	}

	log.Println("Creating Tables")
	err = db.CreateTables()
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("reported to Sentry: %s", err)
		return
	}

	log.Println("Starting bot")
	bot.Start()

}
