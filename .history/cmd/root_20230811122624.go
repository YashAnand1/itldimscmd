/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func itldims(cmd *cobra.Command, args []string) { // Changed the function signature
	response, err := http.Get("http://localhost:8181/servers/")
	if err != nil {
		log.Fatalf("Failed to connect to the etcd API.")
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Connected to API. Interaction with etcd can be done.")
	}
}

var itldimsCmd = &cobra.Command{ // Renamed the variable to itldimsCmd
	Use:   "itldims",
	Short: "Interact with the etcd API",
	Long:  "A command-line tool to interact with the etcd API and check connection",
	Run:   itldims,
}

func init() {
	itldims.AddCommand(get) // Use rootCmd.AddCommand() to add the command
}

func Execute() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
