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

var (
	itldims = &cobra.Command{ //Command will check if connection with API URL is set
		Use:   "itldims",
		Short: "Interact with the etcd API",
		Long:  "A command-line tool to interact with the etcd API and check connection",
		Run:   "itldims",
	}
)

func itldims() (cmd *cobra.Command, args []string) { //here it is defined that when the itldims command is run, it will take a single argument
	response, err := http.Get("http://localhost:8181/servers/") //The response from API is stored in response & error variable
	if err != nil {                                             //if error variable is filled, then the following will be logged
		log.Fatalf("Failed to connect to the etcd API.")
	}

	defer response.Body.Close() //The itldims command & localhost:8181 will be then closed

	if response.StatusCode == http.StatusOK {
		fmt.Println("Connected to API. Interaction with etcd can be done.") //However if the statuscode is ok and connection is set then, message displayed
	}
}

func init() {
	itldims.AddCommand(get)
}

func Execute() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
