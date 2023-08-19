package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "/servers/VM/10.249.221.22/RAM"
	stringSlice := strings.Split(s, "/")
	// SERVERS := stringSlice[1]
	// TYPE := stringSlice[2]
	IP := stringSlice[2]
	// ATTRIBUTE := stringSlice[4]

	fmt.Printf("%v\n", IP /*SERVERS, TYPE, ATTRIBUTE*/)
}
