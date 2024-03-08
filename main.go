package main

import (
	"fmt"
	"strings"
)

func main() {
	/*Bad boy: "SSSRRRRS", "RSSRR", "SRRSSR"
	  Good boy: "SRSSRRR", "SSRR"*/
	var shotAndRevenge string
	fmt.Print("Shot and Revenge: ")
	fmt.Scanf("%1000000s", &shotAndRevenge)
	if strings.HasPrefix(shotAndRevenge, "R") ||
		strings.HasSuffix(shotAndRevenge, "S") ||
		strings.Count(shotAndRevenge, "S") > strings.Count(shotAndRevenge, "R") {
		fmt.Println("Bad boy")
	} else if strings.Count(shotAndRevenge, "S") <= strings.Count(shotAndRevenge, "R") {
		fmt.Println("Good boy")
	} else {
		fmt.Println("Invalid value")
	}
}
