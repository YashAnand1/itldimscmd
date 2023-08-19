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
var get = &cobra.Command{
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using the following combinations:
- 'itldims get <KeyComponent> <KeyComponent/Value>'
- 'itldims get <KeyComponent/Value>'.
s`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

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

			if len(args) == 1 || strings.Contains(args[0], ".") { //if there is only one argument and it consists of "." or ServerIP then
				fmt.Printf("%s:\n%s\n", serverAtr, value) //only display the sererattribute and value

			} else if len(args) == 1 && !strings.Contains(args[0], ".") && strings.Contains(key, args[0]) || strings.Contains(value, args[0]){ //if there is only one argument and it DOESNOT conist of "." or serverIP then
				 //also if the first argument contain key or value as first argument
					fmt.Printf("Server IP: %s\n%s:%s\n", serverIP, serverAtr, value)
				}

			} else {
				if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
					if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
						continue
					}
					fmt.Printf("Server IP: %s\n%s:%s\n", serverIP, serverAtr, value)
				}
			}
		}
	},
}

func init() {
	get.AddCommand(attributes)
	get.AddCommand(types)
	get.AddCommand(servers)
}
