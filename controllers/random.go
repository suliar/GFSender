package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func (c Controller) SendQuotes(from string, to []string, message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Sending Random Quotes")

		err := c.controller.SendQuotes(from, to, message)
		if err != nil {
			log.Fatalf("cannot send SMS %v", err)
		}
	}
}
