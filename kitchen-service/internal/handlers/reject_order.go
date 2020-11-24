package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
)

type RejectTicketConfig struct {
	RejectTicket usecase.RejectTicket
}

type RejectTicketRequest struct {
	OrderId string `json:"order_id"`
}

func NewRejectTicket(config *RejectTicketConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RejectTicketRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.RejectTicket.Execute(request.OrderId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
