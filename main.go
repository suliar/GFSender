package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/suliar/GFSender/controllers"
	quote "github.com/suliar/GFSender/quotes"
	"log"
	"net/http"
	"os"
)

var (
	accountSid string
	token      string
	fromMobile string
	GFMobile   string
	SMobile    string
)

func main() {
	for k, v := range map[string]*string{
		"ACCOUNT_SID": &accountSid,
		"TOKEN":       &token,
		"FROM_MOBILE": &fromMobile,
		"GF_MOBILE":   &GFMobile,
		"S_MOBILE":    &SMobile,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			_ = fmt.Errorf("missing env variables: %s", k)
		}
	}

	controller := controllers.Controller{}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", controller.SendQuotes(quote.NewTwilio(
		fromMobile,
		GFMobile,
		accountSid,
		token))).Methods("GET")

	router.HandleFunc("/dailyBible", controller.SendBibleVerses(quote.NewTwilio(
		fromMobile,
		GFMobile,
		accountSid,
		token))).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
