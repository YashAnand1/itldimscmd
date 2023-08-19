/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using 'itldims get <input1> <input2>' or 'itldims get <input1>'.

Command combinations that can be utilised:
- itldims get <Servers>                | Displays all the running Servers with their Server IPs | Working
- itldims get <Server IP>              | Displays all Attribute values of a specific Server IP | Working
- itldims get <Server Type>            | Displays all Attribute values of a specific Server Type | Working
- itldims get <Attribute>              | Display Servers with a specific Attribute              | Working
- itldims get <Server IP> <Attribute> | Display specific Attribute values of a Server Type     | Working
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		// For itldims get [IP/TYPE/ATTRIBUTE/VALUE] [IP/TYPE/ATTRIBUTE/VALUE]
		// IP ATTRIBUTE, IP TYPE, TYPE ATTRIBUTE the kind
		for key, value := range parseKeyValuePairs(data) {
			if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
				strings.Contains(value, "{") || strings.Contains(value, "}") {
				continue
			}

			if !strings.Contains(key, "data") && (strings.Contains(key, args[0]) || strings.Contains(value, args[0])) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}

				ATs := make(map[string]string)

				splitKey := strings.Split(key, "/")
				serverAtr := splitKey[4]
				serverIP := splitKey[3]
				ATs[serverAtr] = serverAtr

				if len(args) == 1 || strings.Contains(args[0], ".") {
					fmt.Printf("%s:%s\n", serverAtr, value)
				} else {
					fmt.Printf("Server IP: %s\n%s:%s\n", serverIP, serverAtr, value)
				}
			}
		}
	},
}

var get = &cobra.Command {
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: "Data retrieval can be done using 'itldims get <input1> <input2>' or 'itldims get <input1>'."

	data, err := fetchDataFromAPI()
	if err != nil {
		log.Fatalf("Failed to fetch data from the etcd API: %v", err)
	}

	if len(args) == 1 && (args[0] == "servers" || args[0] == "types" || args[0] == "attributes") {
		IPs := make(map[string]string)
		STs := make(map[string]string)
		ATs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverIP := splitKey[3]
			serverType := splitKey[2]
			serverAtr := splitKey[4]

			IPs[serverIP] = serverIP
			STs[serverType] = serverType
			ATs[serverAtr] = serverAtr
		}

		for serverIP := range IPs {
			if args[0] == "servers" {
				fmt.Printf("%s\n", serverIP)
				fmt.Printf("----------------------------\n")
			}
		},
	}
}

		for serverType := range STs {
			if args[0] == "types" {
				fmt.Printf("%s\n", serverType)
				fmt.Printf("----------------------------\n")
			}
		}

		for serverAtr := range ATs {
			if args[0] == "attributes" {
				fmt.Printf("%s\n", serverAtr)
				fmt.Printf("----------------------------\n")
			}
		}

		return
	}

// ... (other code)

func fetchDataFromAPI() (string, error) {
	response, err := http.Get("http://localhost:8181/servers/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func parseKeyValuePairs(data string) map[string]string {
	result := make(map[string]string)

	keyValuePairs := strings.Split(data, "Key:")

	for _, kv := range keyValuePairs {
		lines := strings.Split(kv, "Value:")
		if len(lines) == 2 {
			key := strings.TrimSpace(lines[0])
			value := strings.TrimSpace(lines[1])
			result[key] = value
		}
	}
	return result
}

func init() {
    get.AddCommand(serversCmd)
    getCmd.AddCommand(typesCmd)
    getCmd.AddCommand(attributesCmd)
}
