package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (c Controller) SendQuotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Sending Random Quotes")

		_, err := c.controller.SendQuotes("")
		if err != nil {
			log.Fatalf("cannot send SMS %v", err)
		}
	}
}
