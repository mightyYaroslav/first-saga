package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mightyYaroslav/first-saga/consumer-service/internal/handlers"
	"github.com/mightyYaroslav/first-saga/consumer-service/internal/repository"
	"github.com/mightyYaroslav/first-saga/core/flags"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var serverFlags *flags.Flags

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the consumer service server",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + serverFlags.GetString("db-host") + ":" + serverFlags.GetString("db-port")))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		consumerRepository := repository.NewConsumerRepositoryAdapter(&repository.ConsumerRepositoryAdapterConfig{MongoClient: client})
		r := mux.NewRouter()
		r.HandleFunc(
			"/consumer/verify",
			handlers.NewVerify(&handlers.VerifyConfig{Repository: consumerRepository}),
		).Methods("GET")
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

	serverFlags.RegisterInt("port", "", 8080, "Port of the server", "PORT")
	serverFlags.RegisterString("host", "", "0.0.0.0", "Host of the server", "HOST")

	serverFlags.RegisterString("db-host", "", "mongo", "Host of the database", "DB_HOST")
	serverFlags.RegisterInt("db-port", "", 27017, "Port of the database", "DB_PORT")
}
