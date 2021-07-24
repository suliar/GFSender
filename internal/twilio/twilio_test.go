package twilio_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sfreiberg/gotwilio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/suliar/GFSender/internal/twilio"
	twilio_mock "github.com/suliar/GFSender/mocktwilio/mock/sms"
)

func TestNewTwilio(t *testing.T) {
	t.Run("should return error given from field is missing", func(t *testing.T) {
		twi, err := twilio.NewTwilio("", nil, nil)
		require.Error(t, err)

		assert.Empty(t, twi)

		var e twilio.InvalidParamError
		assert.True(t, errors.As(err, &e))
		assert.Equal(t, "missing from field", e.Parameter)
	})

	t.Run("should return error given to field is missing", func(t *testing.T) {
		twi, err := twilio.NewTwilio("from", nil, nil)
		require.Error(t, err)

		assert.Empty(t, twi)

		var e twilio.InvalidParamError
		assert.True(t, errors.As(err, &e))
		assert.Equal(t, "missing To field", e.Parameter)
	})

	t.Run("should return error given client field is missing", func(t *testing.T) {
		twi, err := twilio.NewTwilio("from", []string{"foo", "bar"}, nil)
		require.Error(t, err)

		assert.Empty(t, twi)

		var e twilio.InvalidParamError
		assert.True(t, errors.As(err, &e))
		assert.Equal(t, "twilio client", e.Parameter)
	})

	t.Run("should return twilio client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTwilio := twilio_mock.NewMockClient(ctrl)

		twi, err := twilio.NewTwilio("from", []string{"foo", "bar"}, mockTwilio)
		require.NoError(t, err)

		assert.NotNil(t, twi)
	})
}

func TestTwilio_SendQuote(t *testing.T) {
	t.Run("should successfully send quotes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const (
			someMessage = "some-messages"
			from        = "from"
		)

		to := []string{"foo", "bar"}
		mockTwilio := twilio_mock.NewMockClient(ctrl)

		mockTwilio.
			EXPECT().
			SendSMS(
				from, gomock.Any(), someMessage, "", "").
			Return(&gotwilio.SmsResponse{},
				&gotwilio.Exception{}, nil).AnyTimes()

		twi, err := twilio.NewTwilio(from, to, mockTwilio)
		require.NoError(t, err)

		err = twi.SendQuote(someMessage)
		require.NoError(t, err)
	})

	t.Run("should return error given issue sending quotes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const (
			someMessage = "some-messages"
			from        = "from"
		)

		to := []string{"foo", "bar"}
		someError := errors.New("some-error")

		mockTwilio := twilio_mock.NewMockClient(ctrl)

		mockTwilio.
			EXPECT().
			SendSMS(
				from, gomock.Any(), someMessage, "", "").
			Return(nil, nil, someError)

		twi, err := twilio.NewTwilio(from, to, mockTwilio)
		require.NoError(t, err)

		err = twi.SendQuote(someMessage)
		require.Error(t, err)
		assert.True(t, errors.Is(err, someError))
	})
}
