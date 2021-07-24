package transport

import (
	"context"
	"net/http"
)

type Bibler interface {
	BibleQuote(ctx context.Context) error
}

type bibleHandler struct {
	bible Bibler
}

func New(bible Bibler) bibleHandler {
	return bibleHandler{bible: bible}
}

func (b bibleHandler) BibleSender(w http.ResponseWriter, r *http.Request) {
	if err := b.bible.BibleQuote(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err := w.Write([]byte("Quote sent"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
