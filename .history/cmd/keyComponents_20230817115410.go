// /*
// Copyright © 2023 NAME HERE <EMAIL ADDRESS>
// */

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var servers = &cobra.Command{ // := only works within functions
	Use:   "servers",
	Short: "Retrieves running servers",
	Long:  "Helps find runing servers by displaying their server IPs",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			fmt.Printf("Could not connect to API. Error: %v", err)
		}
		//map is a built-in data structure that allows you to store key-value pairs. In this case, IPs is a map where both keys and values are of type string.
		IPs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverIP := splitKey[3]
			serverType := splitKey[2]
			IPs[serverIP] = serverType
		}

		for serverIP, serverType := range IPs {
			fmt.Printf("%s:%s\n----------------------------\n", serverIP, serverType)
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
