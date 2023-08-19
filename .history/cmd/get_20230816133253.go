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
	Long: `Data retrieval can be done using 
	- 'itldims get <KeyComponent> <KeyComponent/Value>'
	- 'itldims get <KeyComponent/Value>'.
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}
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
					fmt.Printf("%s:\n%s\n", serverAtr, value)
				} else {
					fmt.Printf("Server IP: %s\n%s:%s\n", serverIP, serverAtr, value)
				}
			}
		}
	},
}

var servers = &cobra.Command{
	Use:   "servers",
	Short: "Displays all the running Servers with their Server IPs",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		IPs := make(map[string]string)

		for key := range parseKeyValuePairs(data) {
			splitKey := strings.Split(key, "/")
			serverIP := splitKey[3]
			IPs[serverIP] = serverIP
		}

		for serverIP := range IPs {
			fmt.Printf("%s\n", serverIP)
			fmt.Printf("----------------------------\n")
		}
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
	get.AddCommand(attributes)
	get.AddCommand(types)
	get.AddCommand(servers)
}
