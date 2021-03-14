package quotes

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

type SendSMSChecker interface {
	SendQuotes(from string, to []string, message string) error
}

type TwilioClient interface {
	SendSMS(from, to, body, statusCallback, applicationSid string) (smsResponse *gotwilio.SmsResponse, exception *gotwilio.Exception, err error)
}

type twilio struct {
	twilio TwilioClient
}

func NewTwilio(client TwilioClient) SendSMSChecker {
	return twilio{
		twilio: client,
	}

}

func (t twilio) SendQuotes(from string, to []string, message string) error {
	if message == "" {
		message = RandomQuote()
	}

	for _, v := range to {
		_, _, err := t.twilio.SendSMS(
			from,
			v,
			message,
			"",
			"")
		if err != nil {
			return fmt.Errorf("send_sms: %w", err)
		}
	}

	return nil
}
