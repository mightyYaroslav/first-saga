package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mightyYaroslav/first-saga/accounting-service/internal/handlers"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the accounting service server",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.HandleFunc("/authorize_credit_card", handlers.NewAuthorizeCreditCard()).Methods("POST")
		http.Handle("/", r)

		// TODO: move it to env
		err := http.ListenAndServe("localhost:8080", r)
		if err != nil {
			log.Fatalf("Could not server the HTTP server as %s:%s", "localhost", "8080")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
