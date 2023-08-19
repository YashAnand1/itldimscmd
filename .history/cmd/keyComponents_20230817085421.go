/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var servers = &cobra.Command{
	Use:   "servers",
	Short: "Displays all the running Servers with their Server IPs",
	Run: func(cmd *cobra.Command, args []string) { // fetch data from API, parse key-value pairs, and display server IPs.
		data, err := fetchDataFromAPI() // function called and retrieved data stored into data variable
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		// IPs := make(map[string]string)										// Map created , this also removes duplicacy of the strings

		for key := range parseKeyValuePairs(data) { // the parsed data is looped and using "/" delimter, key divided into parts
			splitKey := strings.Split(key, "/")
			serverIP := splitKey[3] // the serverIP is the third index of the key
			serverIP = serverIP     // the third index is then saved to the IPs map
		}

		// for serverIP := range serverIP {
		fmt.Printf("%s\n", serverIP)
		fmt.Printf("----------------------------\n")

	},
}

var types = &cobra.Command{
	Use:   "types",
	Short: "Displays all Attribute values of a specific Server Type",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		STs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverType := splitKey[2]
			STs[serverType] = serverType
		}

		for serverType := range STs {
			fmt.Printf("%s\n", serverType)
			fmt.Printf("----------------------------\n")
		}
	},
}

var attributes = &cobra.Command{
	Use:   "attributes",
	Short: "Display Servers with a specific Attribute",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		ATs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverAtr := splitKey[4]
			ATs[serverAtr] = serverAtr
		}

		for serverAtr := range ATs {
			fmt.Printf("%s\n", serverAtr)
			fmt.Printf("----------------------------\n")
		}
	},
}
