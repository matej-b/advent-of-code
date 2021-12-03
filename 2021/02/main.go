package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontal, depth, aim int
}

func main() {
	position := Position{horizontal: 0, depth: 0, aim: 0}
	positionChanges := readInput()
	position = applyPositionModifiers(positionChanges, position)

	// What do you get if you multiply your final horizontal position by your final depth?
	fmt.Println("Final value calculated from horizontal position & depth = ", position.horizontal*position.depth)
}

func applyPositionModifiers(positionChanges []string, position Position) Position {
	for line := 0; line < len(positionChanges); line++ {
		modifier, value := splitLine(positionChanges[line])

		if modifier == "forward" {
			position.horizontal = position.horizontal + value

			if position.aim != 0 {
				position.depth = position.depth + (value * position.aim)
			}
		} else if modifier == "up" {
			// going up decreases the depth
			position.aim = position.aim - value
		} else {
			// going down increases the depth
			position.aim = position.aim + value
		}
	}

	return position
}

// split the line so we get a
// modifier - string - forward, up, down
// value - int
func splitLine(positionChange string) (string, int) {
	temp := strings.Split(positionChange, " ")
	modifier := temp[0]
	value, err := strconv.Atoi(temp[1])

	if err != nil {
		log.Fatal(err)
	}

	return modifier, value
}

// read input file into a string array
func readInput() []string {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return lines
}
