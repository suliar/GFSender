package quotes

import (
	"github.com/sfreiberg/gotwilio"
)

type SendSMSChecker interface {
	SendQuotes(message string) (bool, error)
}

type Twilio struct {
	from       string
	to         string
	accountSid string
	authToken  string
}

func NewTwilio(from string, to string, accountSid string, authToken string) Twilio {
	return Twilio{
		from:       from,
		to:         to,
		authToken:  authToken,
		accountSid: accountSid,
	}

}

func (t *Twilio) SendQuotes(message string) (bool, error) {
	if message == "" {
		message = RandomQuote()
	}

	twilioClient := gotwilio.NewTwilioClient(t.accountSid, t.authToken)
	_, _, err := twilioClient.SendSMS(t.from,
		t.to,
		message,
		"",
		"")
	if err != nil {
		return false, err
	}

	return true, nil
}
