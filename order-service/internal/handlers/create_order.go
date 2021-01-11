package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
)

type CreateOrderConfig struct {
	CreateOrder usecase.CreateOrder
}

type CreateOrderRequest struct {
	TicketId    string `json:"ticket_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TotalPrice  int    `json:"total_price"`
	Status      string `json:"status"`
}

func NewCreateOrder(config *CreateOrderConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateOrderRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = config.CreateOrder.Execute(
			&usecase.CreateOrderParams{
				TicketId:    request.TicketId,
				Title:       request.Title,
				Description: request.Description,
				Status:      request.Status,
				TotalPrice:  request.TotalPrice,
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
