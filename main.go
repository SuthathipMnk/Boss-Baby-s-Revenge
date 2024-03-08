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
	fmt.Scanf("%s", &shotAndRevenge)
	if strings.HasPrefix(shotAndRevenge, "R") ||
		strings.HasSuffix(shotAndRevenge, "S") ||
		strings.Count(shotAndRevenge, "S") > strings.Count(shotAndRevenge, "R") {
		fmt.Println("Bad boy")
	} else {
		fmt.Println("Good boy")
	}
}
