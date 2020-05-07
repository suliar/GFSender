package quotes

import (
	"github.com/sfreiberg/gotwilio"
)

type SendSMSChecker interface {
	SendQuotes(message string) (bool, error)
}

type Twilio struct {
	From       string
	To         string
	AccountSid string
	AuthToken  string
}

func NewTwilio(from string, to string, accountSid string, authToken string) Twilio {
	return Twilio{
		From:       from,
		To:         to,
		AuthToken:  authToken,
		AccountSid: accountSid,
	}

}

func (t *Twilio) SendQuotes(message string) (bool, error) {
	if message == "" {
		message = RandomQuote()
	}

	twilioClient := gotwilio.NewTwilioClient(t.AccountSid, t.AuthToken)

	_, _, err := twilioClient.SendSMS(t.From,
		t.To,
		message,
		"",
		"")
	if err != nil {
		return false, err
	}

	return true, nil
}
