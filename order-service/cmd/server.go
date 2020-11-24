package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mightyYaroslav/first-saga/order-service/internal/handlers"
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the order service server",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Fill this in
		client, err := mongo.NewClient(options.Client().ApplyURI(""))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		orderRepository := repository.NewOrderRepositoryAdapter(&repository.OrderRepositoryAdapterConfig{MongoClient: client})
		createOrder := usecase.NewCreateOrder(&usecase.CreateOrderConfig{OrderRepository: orderRepository})
		approveOrder := usecase.NewApproveOrder(&usecase.ApproveOrderConfig{OrderRepository: orderRepository})
		rejectOrder := usecase.NewRejectOrder(&usecase.RejectOrderConfig{OrderRepository: orderRepository})

		r := mux.NewRouter()
		r.HandleFunc("/order", handlers.NewCreateOrder(&handlers.CreateOrderConfig{CreateOrder: createOrder})).Methods("POST")
		r.HandleFunc("/order/approved", handlers.NewApproveOrder(&handlers.ApproveOrderConfig{ApproveOrder: approveOrder})).Methods("PUT")
		r.HandleFunc("/order/rejected", handlers.NewRejectOrder(&handlers.RejectOrderConfig{RejectOrder: rejectOrder})).Methods("PUT")
		http.Handle("/", r)

		// TODO: move it to env
		err = http.ListenAndServe("localhost:8080", r)
		if err != nil {
			log.Fatalf("Could not server the HTTP server as %s:%s", "localhost", "8080")
		}

		defer client.Disconnect(ctx)
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
