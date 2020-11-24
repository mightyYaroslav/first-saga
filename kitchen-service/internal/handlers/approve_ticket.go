package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
)

type ApproveTicketConfig struct {
	ApproveTicket usecase.ApproveTicket
}

type ApproveTicketRequest struct {
	OrderId string `json:"order_id"`
}

func NewApproveTicket(config *ApproveTicketConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request ApproveTicketRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.ApproveTicket.Execute(request.OrderId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
