package bible_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/suliar/GFSender/internal/bible"
	bible_mock "github.com/suliar/GFSender/mockbible/mock/bibl"
)

func TestService_BibleQuote(t *testing.T) {
	t.Run("should successfully send bible quote", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			ctx          = context.Background()
			message      = "some-message"
			mockQuoter   = bible_mock.NewMockQuoter(ctrl)
			mockRandomer = bible_mock.NewMockRandomer(ctrl)
		)

		mockRandomer.EXPECT().RandomBibleVerses(ctx).Return(message, nil)
		mockQuoter.EXPECT().SendQuote(message).Return(nil)

		s, err := bible.New(mockQuoter, mockRandomer)
		require.NoError(t, err)

		err = s.BibleQuote(ctx)
		require.NoError(t, err)
	})

	t.Run("should return error given issue getting random verses", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			someError    = errors.New("some-error")
			ctx          = context.Background()
			mockQuoter   = bible_mock.NewMockQuoter(ctrl)
			mockRandomer = bible_mock.NewMockRandomer(ctrl)
		)

		mockRandomer.EXPECT().RandomBibleVerses(ctx).Return("", someError)

		s, err := bible.New(mockQuoter, mockRandomer)
		require.NoError(t, err)

		err = s.BibleQuote(ctx)
		require.Error(t, err)
		assert.True(t, errors.Is(err, someError))
	})

	t.Run("should return error given issue sending quote", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			ctx       = context.Background()
			someError = errors.New("some-error")
			message   = "some-message"

			mockQuoter   = bible_mock.NewMockQuoter(ctrl)
			mockRandomer = bible_mock.NewMockRandomer(ctrl)
		)

		mockRandomer.EXPECT().RandomBibleVerses(ctx).Return(message, nil)
		mockQuoter.EXPECT().SendQuote(message).Return(someError)

		s, err := bible.New(mockQuoter, mockRandomer)
		require.NoError(t, err)

		err = s.BibleQuote(ctx)
		require.Error(t, err)
		assert.True(t, errors.Is(err, someError))
	})
}
