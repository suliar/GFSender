package controllers

import (
	"fmt"
	quote "github.com/suliar/GFSender/quotes"
	"log"
	"net/http"
)

type Controller struct {
}

func (c Controller) SendBibleVerses(q quote.Twilio) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		bibleVerse, err := quote.RandomBibleVerses()
		if err != nil {
			log.Fatal(err)
		}

		twiClient := quote.NewTwilio(q.From,
			q.To,
			q.AccountSid,
			q.AuthToken)

		_, err = twiClient.SendQuotes(bibleVerse)
		if err != nil {
			log.Fatalf("cannot send SMS %v", err)
		}
		_, _ = fmt.Fprint(w, "Sending Bible Ve...")
	}
}
