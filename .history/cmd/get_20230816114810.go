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
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		displayServerIP := false
		if len(args) > 0 && !strings.Contains(args[0], ".") {
			displayServerIP = true
			if len(args) > 1 {
				displayServerIP = false
			}
		}

		for key, value := range parseKeyValuePairs(data) {
			if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
				strings.Contains(value, "{") || strings.Contains(value, "}") {
				continue
			}

			if !strings.Contains(key, "data") && (strings.Contains(key, args[0]) || strings.Contains(value, args[0])) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}

				splitKey := strings.Split(key, "/")
				serverAtr := splitKey[4]
				serverIP := splitKey[3]

				if displayServerIP {
					fmt.Printf("%s:%s\n", serverAtr, value)
					fmt.Printf("Server IP: %s\n", serverIP)
					fmt.Printf("--------------------\n")
				} else {
					fmt.Printf("%s\n", value)
					fmt.Printf("--------------------\n")
				}
			}
		}
	},
}

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
