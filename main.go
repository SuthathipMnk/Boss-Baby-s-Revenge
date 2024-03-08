package main

import (
	"fmt"
	"strings"
)

const shot string = "S"
const revenge string = "R"
const goodBoy string = "Good boy"
const badBoy string = "Bad boy"

var output string = goodBoy

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanf("%1000000s", &input)
	printOutput(input)
}

func printOutput(input string) {
	/*Bad boy: "SSSRRRRS", "RSSRR", "SRRSSR"
	  Good boy: "SRSSRRR", "SSRR"*/
	var numOfShot = strings.Count(input, shot)
	var numOfRevenge = strings.Count(input, revenge)

	if strings.HasPrefix(input, revenge) || strings.HasSuffix(input, shot) || numOfShot > numOfRevenge {
		output = badBoy
	} else if numOfShot == numOfRevenge {
		checkCaseShotEqualRevenge(input)
	}

	fmt.Println("Output: ", output)
}

func checkCaseShotEqualRevenge(input string) {
	var countShots int = 0
	var countRevenges int = 0
	var preCharacter string
	for i := 0; i < len(input); i++ {
		var currentCharacter string = string(input[i])
		if currentCharacter != preCharacter && countShots != 0 && countRevenges != 0 {
			if countShots > countRevenges {
				break
			}

			countShots = 0
			countRevenges = 0
		}

		if currentCharacter == shot {
			countShots++
		} else if currentCharacter == revenge {
			countRevenges++
		}
		preCharacter = string(input[i])
	}
	if countShots > countRevenges {
		output = badBoy
	}
}
