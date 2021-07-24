package quotes

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const bibleEndpoint = "/api/?passage=votd"

type client struct {
	rawURL     string
	httpClient *http.Client
}

func New(rawUrl string, httpClient *http.Client) (*client, error) {
	if rawUrl == "" {
		return nil, errors.New("no url found")
	}

	if httpClient == nil {
		httpClient = &http.Client{Timeout: 5 * time.Second}
	}

	return &client{
		rawURL:     rawUrl,
		httpClient: httpClient,
	}, nil
}

func (c *client) RandomBibleVerses(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s%s", c.rawURL, bibleEndpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected error with status code: %d", res.StatusCode)
	}

	verse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	newVerse := fmt.Sprintf("Word of the day:\n\n %s", verse)

	return newVerse, nil
}
