package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	total := 0
	regSet := regexp.MustCompile(`\; \b`)
	regPull := regexp.MustCompile(`\b, \b`)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		possible := true
		subSets := regSet.Split(line, -1)
		for _, curSet := range subSets {
			if !possible {
				break
			}
			cubeMap := map[string]int{"red": 12, "green": 13, "blue": 14}
			subStrings := regPull.Split(curSet, -1)
			for _, curWord := range subStrings {
				split := strings.Split(curWord, " ")
				number, _ := strconv.Atoi(string(split[0]))
				word := split[1]

				remaining := cubeMap[word]
				if remaining >= number {
					remaining -= number
					cubeMap[word] = remaining
				} else {
					possible = false
					fmt.Printf("\n position %d word %s", i, word)
					break
				}
			}
		}
		if possible {
			total += i + 1
		}
	}
	fmt.Println(total)
}
