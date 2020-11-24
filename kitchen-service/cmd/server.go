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
	Short: "Run the kitchen service server",
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

		ticketRepository := repository.NewTicketRepositoryAdapter(&repository.TicketRepositoryAdapterConfig{MongoClient: client})
		approveTicket := usecase.NewApproveOrder(&usecase.ApproveTicketConfig{TicketRepository: ticketRepository})
		rejectTicket := usecase.NewRejectOrder(&usecase.RejectTicketConfig{TicketRepository: ticketRepository})

		r := mux.NewRouter()
		r.HandleFunc("/ticket/approved", handlers.NewApproveTicket(&handlers.ApproveTicketConfig{ApproveTicket: approveTicket})).Methods("PUT")
		r.HandleFunc("/ticket/rejected", handlers.NewRejectTicket(&handlers.RejectTicketConfig{RejectTicket: rejectTicket})).Methods("PUT")
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
