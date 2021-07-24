package twilio

import (
	"fmt"
	"github.com/sfreiberg/gotwilio"
)

type Client interface {
	SendSMS(from, to, body, statusCallback, applicationSid string) (smsResponse *gotwilio.SmsResponse, exception *gotwilio.Exception, err error)
}

type Twilio struct {
	from   string
	to     []string
	twilio Client
}

func NewTwilio(from string, to []string, client Client) (*Twilio, error) {
	switch {
	case from == "":
		return nil, InvalidParamError{Parameter: "missing from field"}
	case to == nil:
		return nil, InvalidParamError{Parameter: "missing To field"}
	case client == nil:
		return nil, InvalidParamError{Parameter: "twilio client"}
	}

	return &Twilio{
		twilio: client,
		from:   from,
		to:     to,
	}, nil
}

func (t *Twilio) SendQuote(message string) error {
	for _, v := range t.to {
		_, _, err := t.twilio.SendSMS(t.from, v, message, "", "")
		if err != nil {
			return fmt.Errorf("send_sms: %w", err)
		}
	}
	return nil
}
