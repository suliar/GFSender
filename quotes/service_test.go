package quotes_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sfreiberg/gotwilio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sms_mock "github.com/suliar/GFSender/mocktwilio/mock/sms"
	"github.com/suliar/GFSender/quotes"
)

type testError string

func (e testError) Error() string { return string(e) }

func TestTwilio_SendQuotes(t *testing.T) {
	t.Run("should return error sending sms", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		const (
			fromMessage = "from-message"
			toMessage   = "to-message"
			someError   = testError("some-error")
		)

		twiolioClient := sms_mock.NewMockTwilioClient(ctrl)

		twiolioClient.
			EXPECT().SendSMS(
			fromMessage,
			toMessage,
			gomock.Any(),
			"",
			"").
			Return(
				nil,
				nil,
				someError,
			)

		smsChecker := quotes.NewTwilio(twiolioClient)

		err := smsChecker.SendQuotes(fromMessage, []string{toMessage}, "")

		require.Error(t, err)

		assert.True(t, errors.Is(err, someError))
	})

	t.Run("should add quotes given quotes are missing from param", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		const (
			fromMessage = "from-message"
			toMessage   = "to-message"
		)

		twiolioClient := sms_mock.NewMockTwilioClient(ctrl)

		twiolioClient.
			EXPECT().SendSMS(
			fromMessage,
			toMessage,
			gomock.Any(),
			"",
			"").
			Return(
				&gotwilio.SmsResponse{},
				&gotwilio.Exception{},
				nil,
			)

		smsChecker := quotes.NewTwilio(twiolioClient)

		err := smsChecker.SendQuotes(fromMessage, []string{toMessage}, "")

		require.NoError(t, err)
	})

	t.Run("send quotes", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		const (
			fromMessage = "from-message"
			toMessage   = "to-message"
			someMessage = "some-message"
		)

		twiolioClient := sms_mock.NewMockTwilioClient(ctrl)

		twiolioClient.
			EXPECT().SendSMS(
			fromMessage,
			toMessage,
			someMessage,
			"",
			"").
			Return(
				&gotwilio.SmsResponse{},
				&gotwilio.Exception{},
				nil,
			)

		smsChecker := quotes.NewTwilio(twiolioClient)

		err := smsChecker.SendQuotes(fromMessage, []string{toMessage}, someMessage)

		require.NoError(t, err)
	})
}
