package cmd

import (
	"fmt"
	"io"       // i/o operations
	"net/http" //
	"strings"
)

func fetchDataFromAPI() (string, error) { //returns string and error
	response, err := http.Get("http://localhost:8181/servers/") //Get request sent to the API URL for fetching data
	if err != nil {
		fmt.Printf("%s", err)
	}

	data, err := io.ReadAll(response.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}

	return string(data), nil //returns the fetched data as a string
}

func parseKeyValuePairs(data string) map[string]string { //string as input and returns a map of strings.
	result := make(map[string]string) //KeyValue pairse to be stored here

	keyValuePairs := strings.Split(data, "Key:")

	for _, kv := range keyValuePairs { //Each keyvalue is gone through the keyValuePairs

		lines := strings.Split(kv, "Value:") //data split into keyvaluepairs is split into kv
		if len(lines) == 2 {                 //if split created 2 lines then key value were split successfuly

			key := strings.TrimSpace(lines[0])
			value := strings.TrimSpace(lines[1])
			result[key] = value
		}
	}

	return result
}
