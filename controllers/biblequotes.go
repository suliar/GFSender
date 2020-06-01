package controllers

import (
	"fmt"
	quote "github.com/suliar/GFSender/quotes"
	"log"
	"net/http"
)

type Controller struct {
	controller quote.SendSMSChecker
}

func New(controller quote.SendSMSChecker) *Controller {
	return &Controller{controller: controller}
}

func (c Controller) SendBibleVerses() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		bibleVerse, err := quote.RandomBibleVerses()
		if err != nil {
			log.Fatal(err)
		}

		_, err = c.controller.SendQuotes(bibleVerse)
		if err != nil {
			log.Fatalf("cannot send SMS %v", err)
		}
		_, _ = fmt.Fprint(w, "Sending Bible Ve...")
	}
}
