package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra" // Import the Cobra library for CLI functionality
)

// Define the 'servers' command
var servers = &cobra.Command{
	Use:   "servers",                                // Command name used in the CLI
	Short: "all servers displayed",                  // Short description of the command's purpose
	Long:  "all servers displayed in a listed form", // Longer description of the command
	Run: func(cmd *cobra.Command, args []string) { // Function to run when the command is executed
		data, err := fetchDataFromAPI() // Fetch data from the API
		if err != nil {                 // Check for API connection error
			fmt.Printf("Unable to connect due to %s", err) // Print the error message
		}

		IPs := make(map[string]string) // Create a map to store server IPs

		// Loop through keys obtained from parsing data
		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/") // Split the key using '/'
			serverIP := splitKey[3]             // Extract the server IP from the split key
			IPs[serverIP] = serverIP            // Store the server IP in the map
		}

		// Loop through stored server IPs and print them
		for serverIP := range IPs {
			fmt.Printf("%s\n----------------------------\n", serverIP) // Print server IP
		}
	},
}

// Define the 'Types' command
var Types = &cobra.Command{
	Use:   "types",                                                     // Command name used in the CLI
	Short: "For finding running server types",                          // Short description of the command's purpose
	Long:  "For finding the running server types from all the servers", // Longer description of the command
	Run: func(cmd *cobra.Command, args []string) { // Function to run when the command is executed
		data, err := fetchDataFromAPI() // Fetch data from the API
		if err != nil {                 // Check for API connection error
			fmt.Printf("Could not connect to API %v", err) // Print the error message
		}

		Types := make(map[string]string) // Create a map to store server types

		// Loop through keys obtained from parsing data
		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/") // Split the key using '/'
			serverType := splitKey[2]           // Extract the server type from the split key
			Types[serverType] = serverType      // Store the server type in the map
		}

		// Loop through stored server types and print them
		for serverType := range Types {
			fmt.Printf("%s\n----------------------------\n", serverType) // Print server type
		}
	},
}

// Define the 'attributes' command
var attributes = &cobra.Command{
	Use:   "attributes",                               // Command name used in the CLI
	Short: "Displays all attributes",                  // Short description of the command's purpose
	Long:  "Displays all attributes that are running", // Longer description of the command
	Run: func(cmd *cobra.Command, args []string) { // Function to run when the command is executed
		data, err := fetchDataFromAPI() // Fetch data from the API
		if err != nil {                 // Check for API connection error
			fmt.Printf("Could not connect to the API. Error: %v", err) // Print the error message
		}

		ATs := make(map[string]string) // Create a map to store attributes

		// Loop through keys obtained from parsing data
		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/") // Split the key using '/'
			attribute := splitKey[4]            // Extract the attribute from the split key
			ATs[attribute] = attribute          // Store the attribute in the map
		}

		// Loop through stored attributes and print them
		for attribute := range ATs {
			fmt.Printf("%s\n----------------------------\n", attribute) // Print attribute
		}

	},
}
