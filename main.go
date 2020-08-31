package main

import (
	"fmt"
	"os"

	"github.com/stevenandrewcarter/terradex/cmd/terradex"
)

func main() {
	if err := terradex.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
