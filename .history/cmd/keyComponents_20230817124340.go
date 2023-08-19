// /*
// Copyright © 2023 NAME HERE <EMAIL ADDRESS>
// */

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var servers = &cobra.Command{
	Use:   "servers",
	Short: "all servers displayed",
	Long:  "all servers displayed in a listed form",
	Run: func(cmd *cobra.Command, args []strings) {
		data, err := fetchDataFromAPI()
		if err != nil {
			fmt.Printf("Unable to connect due to %s", err)
		}

		IPs := make(map[strings]strings)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/") //delimiter
			serverIP := splitKey[3]
			serverAT := splitKey[4]
			IPs[serverIP] = serverAT
		}

		for serverIP, serverAT := range IPs {
			fmt.Printf("%s: %s", serverIP, serverAT)
		}
	},
}

var Types = &cobra.Command{
	Use:   "types",
	Short: "For finding running server types",
	Long:  "For finding the running server types from all the servers",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			fmt.Printf("Could not connect to API %v", err)
		}

		Types := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverType := splitKey[2]
			Types[serverType] = serverType
		}

		for serverType := range Types {
			fmt.Printf("%s\n----------------------------\n", serverType)
		}
	},
}

var attributes = &cobra.Command{
	Use:   "attributes",
	Short: "Displays all attributes",
	Long:  "Displays all attributes that are running",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()

		if err != nil {
			fmt.Printf("Could not connect to the API. Error: %v", err)
		}

		ATs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			attribute := splitKey[4]
			ATs[attribute] = attribute
		}

		for attribute := range ATs {
			fmt.Printf("%s\n----------------------------\n", attribute)
		}

	},
}
