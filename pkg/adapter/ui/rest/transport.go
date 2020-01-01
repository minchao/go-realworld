package rest

import (
	"context"
	"encoding/json"
	"net/http"
)

func BypassRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	// TODO error
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError) // TODO
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
