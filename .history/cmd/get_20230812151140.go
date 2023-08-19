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
	Args: cobra.RangeArgs(1, 2), //Here it is defined that 1 or 2 arguments are only allowed
	Run: func(cmd *cobra.Command, args []string) { //This function is to be executed when 'get' subcommand is run
		data, err := fetchDataFromAPI() //data of fetchDataFromAPI is stored in data variable.
		if err != nil {
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		// For itldims get servers
		if len(args) == 1 && (args[0] == "servers") {
			IPs := make(map[string]string)

			for key := range parseKeyValuePairs(data) {
				splitKey := strings.Split(key, "/")
				serverIP := splitKey[3]

				IPs[serverIP] = serverIP
			}

			for serverIP := range IPs {
				if args[0] == "servers" {
					fmt.Printf("%s\n", serverIP)
				}
			}

			return
		}

		// For itldims get [IP/TYPE/ATTRIBUTE/VALUE] [IP/TYPE/ATTRIBUTE/VALUE]
		for key, value := range parseKeyValuePairs(data) {
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

// This line defines a function named parseKeyValuePairs that takes a single argument data of type string.
// The purpose of this function is to process the input data and return a map containing key-value pairs.
func parseKeyValuePairs(data string) map[string]string {
	//This line initializes an empty map named result with keys and values of type string. This map will store the extracted key-value pairs from the input data.
	result := make(map[string]string)
	//This line uses the strings.Split function to split the input data into an array of substrings based on the delimiter "Key:".
	//Each substring in the resulting keyValuePairs array will correspond to a key-value pair in the input data.
	keyValuePairs := strings.Split(data, "Key:")
	//This line starts a loop that iterates over each element, denoted as kv, in the keyValuePairs array. Each kv represents a segment that contains key-value pair information.
	for _, kv := range keyValuePairs {
		//Within the loop, this line uses the `strings.Split` function again to split the current `kv` segment into an array of substrings based on the delimiter `"Value:"`.
		//This splitting separates the key and value parts of the key-value pair.
		lines := strings.Split(kv, "Value:")
		//This line checks if the array lines resulting from the previous split contains exactly two elements.
		//If it does, it means that the kv segment indeed represents a valid key-value pair (with a key and a value).
		if len(lines) == 2 {
			//These lines assign the trimmed and cleaned key and value parts to the variables key and value, respectively.
			//The strings.TrimSpace function removes any leading or trailing white spaces from the strings.
			key := string(lines[0])
			value := string(lines[1])
			//This line adds the extracted key-value pair to the result map. The key is used as the map key, and the corresponding value is used as the value associated with that key.
			result[key] = value
		}
	}
	//This line returns the populated result map, which contains all the extracted key-value pairs from the input data.
	return result
}
