package quotes

import (
	"fmt"
	"github.com/sfreiberg/gotwilio"
)

type Twilio struct {
	twilio     gotwilio.Twilio
	token      string
	accountSid string
}

func NewQuoteClient(twilio gotwilio.Twilio, token string, accountSid string) *Twilio {
	return &Twilio{
		twilio:     twilio,
		token:      token,
		accountSid: accountSid,
	}
}

func (t *Twilio) SendQuotes(from, to, message string) (bool, error) {
	if message ==  "" {
		message = RandomQuote()
	}
	_, _, err := t.twilio.SendSMS(from,
		to,
		message,
		"",
		"")
	if err != nil {
		return false, fmt.Errorf("could not send SMS: %w", err)
	}
	return true, nil
}


