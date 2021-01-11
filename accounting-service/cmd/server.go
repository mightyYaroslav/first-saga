package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mightyYaroslav/first-saga/accounting-service/internal/handlers"
	"github.com/mightyYaroslav/first-saga/core/flags"
	"github.com/spf13/cobra"
)

var serverFlags *flags.Flags

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the accounting service server",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.HandleFunc("/authorize_credit_card", handlers.NewAuthorizeCreditCard()).Methods("POST")
		http.Handle("/", r)

		// TODO: move it to env
		err := http.ListenAndServe(serverFlags.GetString("host")+":"+serverFlags.GetString("port"), r)
		if err != nil {
			log.Fatalf("Could not server the HTTP server as %s:%s", "localhost", "8080")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
	serverFlags = flags.New("server", serverCommand)

	serverFlags.RegisterInt("port", "p", 8080, "Port of the server", "PORT")
	serverFlags.RegisterString("host", "h", "0.0.0.0", "Host of the server", "HOST")
}
