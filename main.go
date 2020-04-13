package main

import (
	"github.com/sfreiberg/gotwilio"
	"log"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"This is something special - Gege",
		"I want to show you real love - Gege",
		"Happy Birthday to my favourite person - Gege",
		"Thank you for being the best girlfriend ever - Gege",
		"I want to make sure you're doing this - Gege",
	}

	rand.Seed(time.Now().Unix())
	randomQuotes := quotes[rand.Intn(len(quotes))]

	accountSid := "AC8c863d13e1fb9c6068f86926213a0d52"
	token := "88fb19bbcb95938892188a49378a352a"

	twilioClient := gotwilio.NewTwilioClient(accountSid, token)

	_, _, err := twilioClient.SendSMS("+18646252838",
		"+447948847694",
		randomQuotes,
		"",
		"")
	if err != nil {
		log.Fatalf("cannot send sms: %v", err)
	}





}
