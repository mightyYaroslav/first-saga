package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
)

type ApproveOrderConfig struct {
	ApproveOrder usecase.ApproveOrder
}

type ApproveOrderRequest struct {
	OrderId string `json:"order_id"`
}

func NewApproveOrder(config *ApproveOrderConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request ApproveOrderRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = config.ApproveOrder.Execute(request.OrderId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
