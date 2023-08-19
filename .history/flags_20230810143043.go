package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "/servers/VM/10.249.221.22/RAM"
	stringSlice := strings.Split(s, "/")
	SERVERS := stringSlice[0]
	TYPE := stringSlice[1]
	IP := stringSlice[2]
	ATTRIBUTE := stringSlice[3]

	fmt.Printf("%v\n%v\n%v\n%v", SERVERS, TYPE, IP, ATTRIBUTE)
}
