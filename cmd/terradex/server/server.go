package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stevenandrewcarter/terradex/internal/server"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "terradex",
	Short: "Terradex is a HTTP backend for Terraform state configuration and visualization.",
	Long: `Terradex is a Terraform Support Tool for storing remote state with authentication / authorization models
		   and provides a visualization tool to allow you to see how the deployments are going.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := server.Routes()
		log.Fatal(http.ListenAndServe(":8080", r))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
