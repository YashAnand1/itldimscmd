package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "/servers/VM/10.249.221.22/RAM"
	stringSlice := strings.Split(s, "/")
	SERVERS := stringSlice[1]
	TYPE := stringSlice[2]
	IP := stringSlice[3]
	ATTRIBUTE := stringSlice[4]

	fmt.Printf("%v\n%v\n%v\n%v\n", SERVERS, TYPE, IP, ATTRIBUTE)
}
