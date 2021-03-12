package quotes

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

type SendSMSChecker interface {
	SendQuotes(from, to, message string) error
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

func (t twilio) SendQuotes(from string, to, message string) error {
	if message == "" {
		message = RandomQuote()
	}

	_, _, err := t.twilio.SendSMS(
		from,
		to,
		message,
		"",
		"")
	if err != nil {
		return fmt.Errorf("send_sms: %w", err)
	}

	return nil
}
