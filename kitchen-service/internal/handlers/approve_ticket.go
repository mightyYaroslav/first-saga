package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/usecase"
)

type ApproveTicketConfig struct {
	ApproveTicket usecase.ApproveTicket
}

type ApproveTicketRequest struct {
	TicketId string `json:"ticket_id"`
}

func NewApproveTicket(config *ApproveTicketConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request ApproveTicketRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.ApproveTicket.Execute(request.TicketId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
