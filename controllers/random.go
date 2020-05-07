package controllers

import (
	"fmt"
	quote "github.com/suliar/GFSender/quotes"
	"log"
	"net/http"
)

func (c Controller) SendQuotes(q quote.Twilio) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Sending Random Quotes")
		twiClient := quote.NewTwilio(q.From,
			q.To,
			q.AccountSid,
			q.AuthToken)
		_, err := twiClient.SendQuotes("")
		if err != nil {
			log.Fatalf("cannot send SMS %v", err)
		}
	}
}
