package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/usecase"
)

type RejectTicketConfig struct {
	RejectTicket usecase.RejectTicket
}

type RejectTicketRequest struct {
	TicketId string `json:"ticket_id"`
}

func NewRejectTicket(config *RejectTicketConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RejectTicketRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.RejectTicket.Execute(request.TicketId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
