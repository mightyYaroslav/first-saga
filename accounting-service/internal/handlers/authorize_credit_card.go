package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type AuthorizeCreditCardRequest struct {
	CreditCardNumber string `json:"credit_card_number"`
}

func NewAuthorizeCreditCard() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request AuthorizeCreditCardRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rnd := rand.Intn(100)

		if rnd > 50 {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			http.Error(w, "Wrong credit card number", http.StatusForbidden)
			return
		}
	}
}
