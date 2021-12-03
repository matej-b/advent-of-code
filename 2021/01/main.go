package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var previousNo = 0
var countLarger = 0
var results []int

func main() {
	// add a --debug flag to provide extra output
	debugPtr := flag.Bool("debug", false, "a bool")
	flag.Parse()

	values := readInput()
	checkValues(values)
	fmt.Println(countLarger)

	// print all results included in count if --debug flag is set
	if *debugPtr {
		fmt.Println(results)
	}
}

func readInput() []string {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	return lines
}

func checkValues(values []string) {
	for i := 0; i < len(values); i++ {
		currentNo, err := strconv.Atoi(values[i])

		if err != nil {
			log.Fatal(err)
		}

		// we have nothing to compare the first line to
		// so set the value and skip to the next
		if previousNo == 0 {
			previousNo = currentNo
			continue
		}

		if previousNo < currentNo {
			results = append(results, currentNo)
			countLarger++
		}

		previousNo = currentNo
	}
}
