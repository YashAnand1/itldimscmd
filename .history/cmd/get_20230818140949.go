package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{ // Cobra command is 'get'
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using the following combinations:
 - 'itldims get <KeyComponent> <KeyComponent/Value>'
 - 'itldims get <KeyComponent/Value>'
`,
	Args: cobra.RangeArgs(1, 2), //number of allowed arguments=1&2
	Run: func(cmd *cobra.Command, args []string) {

		data, err := fetchDataFromAPI() // Call a function to fetch data from the API
		if err != nil {                 // Check if an error occurred
			fmt.Printf("Failed to fetch data from the etcd API: %v", err) // Print an error message
		}

		// Loop through key-value pairs obtained from parsing data
		for key, value := range parseKeyValuePairs(data) {
			// Skip pairs containing specific substrings
			if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
				strings.Contains(value, "{") || strings.Contains(value, "}") {
				continue // Skip this iteration of the loop
			}

			ATs := make(map[string]string)      // Create a map to store attributes
			splitKey := strings.Split(key, "/") // Split the key using '/' delimiter
			serverAtr := splitKey[4]            // Extract the server attribute from the split key
			serverIP := splitKey[3]             // Extract the server IP from the split key
			ATs[serverAtr] = serverAtr          // Store the server attribute in the map

			// Check conditions to decide on printing
			if strings.Contains(args[0], ".") && (strings.Contains(key, args[0]) || strings.Contains(value, args[0])) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue // Skip this iteration of the loop
				}
				fmt.Printf("%s:\n%s\n\n", serverAtr, value) // Print server attribute and value
			} else {
				if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
					if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
						continue // Skip this iteration of the loop
					}
					fmt.Printf("Server IP:\n%s%s:%s\n\n", serverIP, serverAtr, value) // Print server IP, attribute, and value
				}
			}
		}
	},
}

// Initialize the 'get' command and add subcommands
func init() {
	get.AddCommand(attributes) // Add 'attributes' subcommand
	get.AddCommand(Types)      // Add 'Types' subcommand
	get.AddCommand(servers)    // Add 'servers' subcommand
}
