package main

import (
	"fmt"
	"strings"
)
func main (){
	var s string = "/servers/VM/10.249.221.22/RAM"
	stringSlice := strings.Split(s, "/")
	SERVERS := stringSlice[0]
	TYPE := stringSlice[1]
	IP := stringSlice [2]
	ATTRIBUTE := stringSlice[3]

	fmt.println("%v\n%v\n%v\n%v", SERVERS, TYPE, IP, ATTRIBUTE)
}

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