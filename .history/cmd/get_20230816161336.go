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

// For itldims get [IP/TYPE/ATTRIBUTE/VALUE] [IP/TYPE/ATTRIBUTE/VALUE]
// IP ATTRIBUTE, IP TYPE, TYPE ATTRIBUTE the kind
Run: func(cmd *cobra.Command, args []string) {
	data, err := fetchDataFromAPI()
	if err != nil {
		log.Fatalf("Failed to fetch data from the etcd API: %v", err)
	}

	firstArg := args[0]
	serverIPRequested := strings.Contains(firstArg, ".")

	for key, value := range parseKeyValuePairs(data) {
		if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
			strings.Contains(value, "{") || strings.Contains(value, "}") {
			continue
		}
	
		ATs := make(map[string]string)
		splitKey := strings.Split(key, "/")
		serverAtr := splitKey[4]
		serverIP := splitKey[3]
		ATs[serverAtr] = serverAtr
	
		if len(args) == 1 && !serverIPRequested {
			if strings.Contains(key, firstArg) || strings.Contains(value, firstArg) {
				fmt.Printf("Server IP: %s\n%s:%s\n----------------------------\n", serverIP, serverAtr, value)
			}
		} else if len(args) == 1 || serverIPRequested {
			fmt.Printf("%s:%s\n----------------------------\n", serverAtr, value)
		} else {
			if strings.Contains(key, firstArg) || strings.Contains(value, firstArg) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}
				fmt.Printf("Server IP: %s\n%s:%s\n----------------------------\n", serverIP, serverAtr, value)
			}
			
		}

	},
}

func init() {
	get.AddCommand(attributes)
	get.AddCommand(types)
	get.AddCommand(servers)
}
