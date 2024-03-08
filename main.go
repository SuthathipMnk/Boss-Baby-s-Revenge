package main

import (
	"fmt"
	"strings"
)

const shot string = "S"
const revenge string = "R"
const goodBoy string = "Good boy"
const badBoy string = "Bad boy"

func main() {
	/*Bad boy: "SSSRRRRS", "RSSRR", "SRRSSR"
	  Good boy: "SRSSRRR", "SSRR"*/
	var shotAndRevenge string
	fmt.Print("Input: ")
	fmt.Scanf("%1000000s", &shotAndRevenge)
	var numOfShot = strings.Count(shotAndRevenge, shot)
	var numOfRevenge = strings.Count(shotAndRevenge, revenge)
	var result string = goodBoy
	if strings.HasPrefix(shotAndRevenge, revenge) ||
		strings.HasSuffix(shotAndRevenge, shot) ||
		numOfShot > numOfRevenge {
		result = badBoy
	} else if numOfShot == numOfRevenge {
		var countShots int = 0
		var countRevenges int = 0
		var preCharacter string
		for i := 0; i < len(shotAndRevenge); i++ {
			var currentCharacter string = string(shotAndRevenge[i])
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
			preCharacter = string(shotAndRevenge[i])
		}
		if countShots > countRevenges {
			result = badBoy
		}
	}
	fmt.Println("Output: ", result)
}
