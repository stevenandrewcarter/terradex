package terradex

import (
	"log"

	"github.com/spf13/cobra"
)

// VersionCmd for getting Version information
var VersionCmd = &cobra.Command{
	Use:   "terradex",
	Short: "Terradex is a HTTP backend for Terraform state configuration and visualization.",
	Long: `Terradex is a Terraform Support Tool for storing remote state with authentication / authorization models
		     and provides a visualization tool to allow you to see how the deployments are going.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf(`Version 0.1.0`)
	},
}
