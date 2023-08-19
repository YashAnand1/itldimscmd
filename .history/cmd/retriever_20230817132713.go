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
		return "", err
	}

	data, err := io.ReadAll(response.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
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
