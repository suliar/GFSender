package main

import (
	"fmt"
	"github.com/gorilla/mux"
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

func sendQuotes(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Sending Random Quotes")
	twiClient := quote.NewTwilio(fromMobile,
		SMobile,
		accountSid,
		token)
	_, err := twiClient.SendQuotes("")
	if err != nil {
		log.Fatalf("cannot send SMS %v", err)
	}
}

func SendRandomBibleVerses(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Sending Bible Verses...")
	bibleVerse, err := quote.RandomBibleVerses()
	if err != nil {
		log.Fatal(err)
	}

	twiClient := quote.NewTwilio(fromMobile,
		GFMobile,
		accountSid,
		token)

	_, err = twiClient.SendQuotes(bibleVerse)
	if err != nil {
		log.Fatalf("cannot send SMS %v", err)
	}
}

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", sendQuotes).Methods("GET")
	router.HandleFunc("/dailyBible", SendRandomBibleVerses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
