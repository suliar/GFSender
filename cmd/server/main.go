package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sfreiberg/gotwilio"

	"github.com/suliar/GFSender/internal/bible"
	quote "github.com/suliar/GFSender/internal/quotes"
	transportHttp "github.com/suliar/GFSender/internal/transport/http"
	"github.com/suliar/GFSender/internal/twilio"
)

var (
	accountSid string
	token      string
	fromMobile string
	GFMobile   string
	SMobile    string
	MMobile    string
	BibleURL   string
)

func main() {
	for k, v := range map[string]*string{
		"ACCOUNT_SID": &accountSid,
		"TOKEN":       &token,
		"FROM_MOBILE": &fromMobile,
		"GF_MOBILE":   &GFMobile,
		"S_MOBILE":    &SMobile,
		"M_MOBILE":    &MMobile,
		"BIBLE_URL":   &BibleURL,
	} {
		var ok bool
		if *v = os.Getenv(k); !ok {
			_ = fmt.Errorf("missing env variables: %s", k)
		}
	}

	numbers := []string{GFMobile, SMobile, MMobile}
	twilioClient := gotwilio.NewTwilioClient(accountSid, token)

	twn, err := twilio.NewTwilio(fromMobile, numbers, twilioClient)
	if err != nil {
		log.Fatal("could not create new twilio")
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	quo, err := quote.New(BibleURL, client)
	if err != nil {
		log.Fatal("could not create new quote")
		return
	}

	ser, err := bible.New(twn, quo)
	if err != nil {
		log.Fatal("could not create new bible")
		return
	}

	handler := transportHttp.New(ser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	router := mux.NewRouter()
	router.HandleFunc("/daily", handler.BibleSender).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
