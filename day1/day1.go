package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
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
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		num1, num2 := 0, 0
		pos1 := 0
		for i, letter := range line {
			number, err := strconv.Atoi(string(letter))
			if err == nil {
				num1 = number
				pos1 = i
				break
			}
		}
		for i := len(line) - 1; i >= pos1; i-- {
			letter := line[i]
			number, err := strconv.Atoi(string(letter))
			if err == nil {
				num2 = number
				break
			}
		}
		if num2 == 0 {
			num2 = num1
		}
		combined, _ := strconv.Atoi(strings.Join([]string{strconv.Itoa(num1), strconv.Itoa(num2)}, ""))
		total += combined
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partTwo() {
	f, err := os.Open("input")
	numWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	var total int
	pos := 1
	for scanner.Scan() {
		line := scanner.Text()
		num1, num2 := 0, 0
		currentWord := ""
		for i := 0; i <= len(line); i++ {
			letter := line[i]
			number, err := strconv.Atoi(string(letter))
			if err == nil {
				num1 = number
				break
			} else {
				currentWord += string(letter)
				found := false
				exactMatch := false
				for j, word := range numWords {
					if word == currentWord {
						num1 = j + 1
						exactMatch = true
						break
					}
					if strings.HasPrefix(word, currentWord) {
						found = true
						break
					}
				}
				if exactMatch {
					break
				}
				if !found {
					//have to set I back to last possible place
					i -= len(currentWord) - 1
					currentWord = ""
				}
			}
		}
		currentWord = ""
		for i := len(line) - 1; i >= 0; i-- {
			letter := line[i]
			number, err := strconv.Atoi(string(letter))
			if err == nil {
				num2 = number
				break
			} else {
				currentWord = string(letter) + currentWord
				found := false
				exactMatch := false
				for j, word := range numWords {
					if word == currentWord {
						num2 = j + 1
						exactMatch = true
						break
					}
					if strings.HasSuffix(word, currentWord) {
						found = true
						break
					}
				}
				if exactMatch {
					break
				}
				if !found {
					i += len(currentWord) - 1
					currentWord = ""
				}
			}
		}

		if num2 == 0 {
			num2 = num1
		}

		combined, _ := strconv.Atoi(strings.Join([]string{strconv.Itoa(num1), strconv.Itoa(num2)}, ""))
		pos++
		total += combined
	}
	fmt.Println("")
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Print("p1: ")
	partOne()
	fmt.Print("p2: ")
	partTwo()
}
