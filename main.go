package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type NameData struct {
	Rank string
	Boy  string
	Girl string
}

func main() {
	// Open the HTML file
	file, err := os.Open("2022babynames.html")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a regular expression pattern
	pattern := regexp.MustCompile(`<td>(\d+)</td> <td>(\w+)</td> <td>(\w+)</td>`)

	// Read the file line by line and load data into a slice of slices
	var nameData [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		if len(matches) == 4 {
			nameData = append(nameData, matches[1:])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Prompt the user for the initial gender
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter gender (boy/girl): ")
	genderInput, _ := reader.ReadString('\n')
	gender := strings.TrimSpace(strings.ToLower(genderInput))

	// Prompt the user for the number of names to choose between
	fmt.Print("Enter the number of names to choose between: ")
	numNamesInput, _ := reader.ReadString('\n')
	numNamesInput = strings.TrimSpace(numNamesInput)
	numNames := len(nameData)

	if numNamesInput != "" {
		numNames, _ = strconv.Atoi(numNamesInput)
		if numNames > len(nameData) || numNames < 1 {
			numNames = len(nameData)
		}
	}

	// Generate and output random names based on user input
	rand.Seed(time.Now().UnixNano())
	for {
		if gender == "boy" {
			fmt.Println("Press enter (or type 'girl')")
		} else if gender == "girl" {
			fmt.Println("Press enter (or type 'boy')")
		} else {
			fmt.Println("Invalid gender choice. Please choose 'boy' or 'girl'.")
			break
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "boy" {
			gender = "boy"
		} else if input == "girl" {
			gender = "girl"
		} else if input == "quit" || input == "q" {
			fmt.Println("Goodbye.")
			break
		}

		if gender == "boy" {
			choose := nameData[rand.Intn(numNames)]
			name := color.New(color.FgBlue).Sprint(choose[1])
			fmt.Printf("Random %s name: %s (rank: %s)\n", gender, name, choose[0])
		} else if gender == "girl" {
			choose := nameData[rand.Intn(numNames)]
			name := color.New(color.FgMagenta).Sprint(choose[2])
			fmt.Printf("Random %s name: %s (rank: %s)\n", gender, name, choose[0])
		} else {
			fmt.Println("Invalid gender choice. Please choose 'boy' or 'girl'.")
			break
		}
	}
}
