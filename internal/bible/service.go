package bible

import (
	"context"
	"fmt"
)

type Quoter interface {
	SendQuote(message string) error
}

type Randomer interface {
	RandomBibleVerses(ctx context.Context) (string, error)
}

type service struct {
	quote Quoter
	rand  Randomer
}

func New(quote Quoter, rand Randomer) (service, error) {
	return service{
		quote: quote,
		rand:  rand,
	}, nil

}

func (s service) BibleQuote(ctx context.Context) error {
	verse, err := s.rand.RandomBibleVerses(ctx)
	if err != nil {
		return fmt.Errorf("error_getting_random_verse: %w", err)
	}

	if err := s.quote.SendQuote(verse); err != nil {
		return fmt.Errorf("error_sending_quote: %w", err)
	}

	return nil
}
