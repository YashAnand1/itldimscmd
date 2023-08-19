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
	Short: "Displays attributes and values for running servers",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		serverAttrs := make(map[string]map[string]string{}

		for key, value := range parseKeyValuePairs(data) {
			if strings.Contains(key, "{") || strings.Contains(value, "{") {
				continue
			}

			serverIP, serverAttr := strings.Split(key, "/")[3], strings.Split(key, "/")[4]

			if _, exists := serverAttrs[serverIP]; !exists {
				serverAttrs[serverIP] = map[string]string{}
			}

			serverAttrs[serverIP][serverAttr] = value
		}

		for serverIP, attrs := range serverAttrs {
			fmt.Printf("Server IP: %s\n", serverIP)
			for attr, value := range attrs {
				fmt.Printf("%s: %s\n", attr, value)
			}
			fmt.Println("----------------------------")
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
