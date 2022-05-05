package coders

import (
	"context"
	"encoding/json"
	"net/http"
	"test_robert_yofio/internal/entity"
)

func DecodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req entity.GenericRequest
	return req, nil
}

func DecodeInvestmentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req entity.InvestmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func RespondWithJSON(_ context.Context, w http.ResponseWriter, payload interface{}) error {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)
	return err
}
