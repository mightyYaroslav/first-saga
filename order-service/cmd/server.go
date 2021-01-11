package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mightyYaroslav/first-saga/core/flags"
	"github.com/mightyYaroslav/first-saga/order-service/internal/handlers"
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
	"github.com/mightyYaroslav/first-saga/order-service/internal/usecase"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var serverFlags *flags.Flags

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the order service server",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Fill this in
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + serverFlags.GetString("db-host") + ":" + serverFlags.GetString("db-port")))
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

		err = http.ListenAndServe(serverFlags.GetString("host")+":"+serverFlags.GetString("port"), r)
		if err != nil {
			log.Fatalf("Could not server the HTTP server as %s:%s", "localhost", "8080")
		}

		defer client.Disconnect(ctx)
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
	serverFlags = flags.New("server", serverCommand)

	serverFlags.RegisterInt("port", "p", 8080, "Port of the server", "PORT")
	serverFlags.RegisterString("host", "h", "0.0.0.0", "Host of the server", "HOST")

	serverFlags.RegisterString("db-host", "db-host", "mongo", "Host of the database", "DB_HOST")
	serverFlags.RegisterInt("db-port", "db-port", 27017, "Port of the database", "DB_PORT")
}
