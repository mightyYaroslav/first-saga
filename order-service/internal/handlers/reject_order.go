package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
)

type RejectOrderConfig struct {
	RejectOrder usecase.RejectOrder
}

type RejectOrderRequest struct {
	OrderId string `json:"order_id"`
}

func NewRejectOrder(config *RejectOrderConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RejectOrderRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.RejectOrder.Execute(request.OrderId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
