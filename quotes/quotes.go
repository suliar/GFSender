package quotes

import (
	"math/rand"
	"time"
)

var quotes = []string{
	"Today is a special day",
	"I want to show you real love",
	"Happy Birthday to my favourite person",
	"Thank you for being the best girlfriend ever",
	"I want to make sure you're doing this",
}

func RandomQuote() string {
	rand.Seed(time.Now().Unix())
	randomQuote := quotes[rand.Intn(len(quotes))]
	return randomQuote
}

