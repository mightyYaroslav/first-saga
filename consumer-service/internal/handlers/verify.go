package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/consumer-service/internal/repository"
)

type VerifyConfig struct {
	Repository repository.Consumer
}

type VerifyRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func NewVerify(config *VerifyConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request VerifyRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.Repository.Verify(
			&repository.ConsumerVerifyParams{
				FirstName: request.FirstName,
				LastName:  request.LastName,
				Phone:     request.Phone,
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
