package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "order-service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(context.Background(), "error in rootCmd.Execute", err.Error())
	}
}
