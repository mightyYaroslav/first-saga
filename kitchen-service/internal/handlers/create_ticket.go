package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/usecase"
)

type CreateTicketConfig struct {
	CreateTicket usecase.CreateTicket
}

type CreateTicketRequest struct {
	Id      string   `json:"id"`
	OrderId string   `json:"order_id"`
	Title   string   `json:"title"`
	Dishes  []string `json:"dishes"`
}

func NewCreateTicket(config *CreateTicketConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateTicketRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.CreateTicket.Execute(
			&usecase.CreateTicketParams{
				Id:      request.Id,
				OrderId: request.OrderId,
				Title:   request.Title,
				Dishes:  request.Dishes,
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
