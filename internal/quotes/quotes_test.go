package quotes_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	quote "github.com/suliar/GFSender/internal/quotes"
)

func TestNew(t *testing.T) {
	t.Run("should return error given issue parsing url", func(t *testing.T) {
		client, err := quote.New("%%%", nil)
		require.Error(t, err)

		assert.Nil(t, client)
	})

	t.Run("should return client", func(t *testing.T) {
		client, err := quote.New("http://something.com", nil)
		require.NoError(t, err)

		assert.NotNil(t, client)
	})
}

func TestClient_RandomBibleVerses(t *testing.T) {
	t.Run("should successfully return a bible verse", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			verse := "some-bible-verse"

			bb, err := json.Marshal(verse)
			require.NoError(t, err)

			_, err = w.Write(bb)
			require.NoError(t, err)
		}))
		defer server.Close()

		client, err := quote.New(server.URL, nil)
		require.NoError(t, err)

		res, err := client.RandomBibleVerses(context.Background())
		require.NoError(t, err)

		assert.NotEmpty(t, res)
	})

	t.Run("should return error given status is not 200", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client, err := quote.New(server.URL, nil)
		require.NoError(t, err)

		res, err := client.RandomBibleVerses(context.Background())
		require.Error(t, err)

		assert.Empty(t, res)
	})
}
