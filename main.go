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
	toMobile   string
)


func sendQuotes( w http.ResponseWriter, r *http.Request ) {
	twiClient := quote.NewTwilio(fromMobile,
		toMobile,
		accountSid,
		token)
	_, err := twiClient.SendQuotes("")
	if err != nil {
		log.Fatalf("cannot send SMS %v", err)
	}

}
func main() {

	for k, v := range map[string]*string{
		"ACCOUNT_SID": &accountSid,
		"TOKEN":       &token,
		"FROM_MOBILE": &fromMobile,
		"TO_MOBILE":   &toMobile,
	} {
		var ok bool
		if *v, ok = os.LookupEnv(k); !ok {
			_ = fmt.Errorf("missing env variables: %s", k)
		}
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/quotes", sendQuotes).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
