package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{ //the command will be used for retrieving data by filtering content of the API Server
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using 'itldims get <input1> <input2>' or 'itldims get <input1>'.

Command combinations that can be utilised:
- itldims get <Servers>          		| Displays all the running Servers with their Server IPs | Working
- itldims get <Server IP>        		| Displays all Attribute values of a specific Server IP | Working
- itldims get <Server Type>        		| Displays all Attribute values of a specific Server Type | Working
- itldims get <Attribute>        		| Display Servers with a specific Attribute
- itldims get <Value>            		| Display Servers with a specific Attribute value
- itldims get <Server Type> <Attribute> | Display specific Attribute values of a Server Type | Working
- itldims get <Server Type> <Value>     | Display Server Types containing a specific value
- itldims get <Value> <Server Type>     | Display Server Types containing a specific value
- itldims get <Attribute> <Server IP>   | Display specific Attribute values of a Server IP
- itldims get <Server IP> <Attribute>   | Display specific Attribute values of a Server IP
- itldims get <Server IP> <Value>       | Display Server IPs containing a specific value
- itldims get <Value> <Server IP>       | Display Server IPs containing a specific value
- itldims get <Server IP> <Server Type> | Display Attribute values of a specific Server
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		if len(args) == 1 && (args[0] == "servers") || (args[0] == "types") || (args[0] == "attributes") { // For itldims get servers
			IPs := make(map[string]string)
			ATs := make(map[string]string)
			STs := make(map[string]string)

			for key := range parseKeyValuePairs(data) {
				splitKey := strings.Split(key, "/")
				serverIP := splitKey[3]
				serverTy := splitKey[2]
				serverAt := splitKey[4]

				IPs[serverIP] = serverIP
				STs(serverTy) = serverTy
				ATs[serverAt] = serverIP

			}

			fmt.Printf("Servers IP 		| Server Type | Attribute\n")
			for serverIP := range IPs {
				if args[0] == "servers" {
					fmt.Printf("%s\n", serverIP)
					fmt.Printf("----------------------------------------\n\n")
				}
			}

			return
		}

		for key, value := range parseKeyValuePairs(data) { // For itldims get [IP/TYPE/ATTRIBUTE/VALUE] [IP/TYPE/ATTRIBUTE/VALUE]
			if strings.Contains(value, "{") || strings.Contains(value, "}") || strings.Contains(key, "data") {
				continue
			}

			if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}

				fmt.Printf("%s:\n%s\n\n", key, value)
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
			key := string(lines[0])
			value := string(lines[1])
			result[key] = value
		}
	}
	return result
}
