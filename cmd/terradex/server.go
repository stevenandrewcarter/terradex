package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stevenandrewcarter/terradex/internal/server"
	"log"
	"net/http"
	"os"
)

var port string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "terradex",
	Short: "Terradex is a HTTP backend for Terraform state configuration and visualization.",
	Long: `Terradex is a Terraform Support Tool for storing remote state with authentication / authorization models
		     and provides a visualization tool to allow you to see how the deployments are going.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf(`Starting Terradex (Version 0.0.1)
 ____  ____  ____  ____   __   ____  ____  _  _ 
(_  _)(  __)(  _ \(  _ \ / _\ (    \(  __)( \/ )
  )(   ) _)  )   / )   //    \ ) D ( ) _)  )  ( 
 (__) (____)(__\_)(__\_)\_/\_/(____/(____)(_/\_)`)
		log.Printf("Loading Routes...")
		r := server.Routes()
		log.Printf("Started and running...")
		log.Printf("Listening on Port %s", port)
		log.Fatal(http.ListenAndServe(":"+port, r))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.terradex.yaml)")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "The Port that Terradex will listen on.")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".terradex")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		// os.Exit(1)
	}
}
