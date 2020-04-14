package main

import (
	"fmt"
	quote "github.com/suliar/GFSender/quotes"
	"log"
	"os"
)

func main() {

	var (
		accountSid string
		token      string
		fromMobile string
		toMobile   string
	)

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

	twiClient := quote.NewTwilio(fromMobile,
		toMobile,
		accountSid,
		token)
	res, err := twiClient.SendQuotes("do something")
	if err != nil {
		log.Fatalf("cannot send SMS %v", err)
	}

	fmt.Println(res)
}
