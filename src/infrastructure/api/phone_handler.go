package api

import (
	"encoding/json"
	"net/http"

	"github.com/ViniciusMartinss/phone-number-handler/src"
	"github.com/ViniciusMartinss/phone-number-handler/src/domain"
)

type apiResponse struct {
	Status bool                   `json:"status"`
	Count  int                    `json:"count"`
	Result []domain.PhoneReturnee `json:"result"`
}

func (s server) phoneListHandler(w http.ResponseWriter, r *http.Request) {
	filters := setFilters(r)

	result := s.phoneHandler.
		List(filters)

	response := apiResponse{
		Status: true,
		Count:  len(result),
		Result: result,
	}

	w.Header().
		Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).
		Encode(response)
}

func setFilters(r *http.Request) map[string]string {
	country := r.URL.
		Query().
		Get(src.COUNTRY)

	state := r.URL.
		Query().
		Get(src.STATE)

	return map[string]string{
		src.COUNTRY: country,
		src.STATE:   state,
	}
}
