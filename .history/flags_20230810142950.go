package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "/My/Name/Is/Yash"
	stringSlice := strings.Split(s, "/")
	My := stringSlice[1]
	Name := stringSlice[2]
	Is := stringSlice[3]
	Yash := stringSlice[4]

	fmt.Printf("%v\n%v\n%v\n%v\n", My, Name, Is, Yash)
}
