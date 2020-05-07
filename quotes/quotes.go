package quotes

import (
	"context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"Today is a special day",
	"I want to show you real love",
	"Happy Birthday to my favourite person",
	"Thank you for being the best girlfriend ever",
	"I want to make sure you're doing this",
	"What would I do without you aye",
	"I'm so proud of you",
	"I want you to know that I love you",
}

func RandomQuote() string {
	rand.Seed(time.Now().Unix())
	randomQuote := quotes[rand.Intn(len(quotes))]
	return randomQuote
}

func RandomBibleVerses() (string, error) {
	req, err := http.NewRequest("GET",
		"http://www.ourmanna.com/verses/api/get?format=text&order=random",
		nil)

	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	verses, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(verses), nil

}
