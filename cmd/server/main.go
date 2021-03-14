package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sfreiberg/gotwilio"

	"github.com/suliar/GFSender/controllers"
	quote "github.com/suliar/GFSender/quotes"
)

var (
	accountSid string
	token      string
	fromMobile string
	GFMobile   string
	SMobile    string
	MMobile    string
)

func main() {
	for k, v := range map[string]*string{
		"ACCOUNT_SID": &accountSid,
		"TOKEN":       &token,
		"FROM_MOBILE": &fromMobile,
		"GF_MOBILE":   &GFMobile,
		"S_MOBILE":    &SMobile,
		"M_MOBILE":    &MMobile,
	} {
		var ok bool
		if *v = os.Getenv(k); !ok {
			_ = fmt.Errorf("missing env variables: %s", k)
		}
	}

	twilioClient := gotwilio.NewTwilioClient(accountSid, token)

	cm := quote.NewTwilio(twilioClient)

	newController := controllers.New(cm)

	router := mux.NewRouter().StrictSlash(true)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	numbers := []string{GFMobile, SMobile, MMobile}

	router.HandleFunc("/daily-bible",
		newController.SendBibleVerses(fromMobile, numbers)).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))

}
