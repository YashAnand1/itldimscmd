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
