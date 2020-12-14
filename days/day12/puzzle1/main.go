package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

//Directions
type Direction struct {
	direction string
	steps     int
}

func parseDirections(input []string) []Direction {

	//create array of struct Direction
	directions := make([]Direction, len(input))
	//split based on alphanumeric
	re := regexp.MustCompile(`[A-Z]`)

	for i, s := range input {

		//first byte is direction
		directions[i].direction = string(s[0])
		//second byte(s) is the steps
		result := re.Split(s, -1)
		directions[i].steps, _ = strconv.Atoi(result[1])

	}
	return directions
}

func parseInputFromFile(path string) []Direction {

	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return parseDirections(lines)
}

func printManhattenDirections(directions []Direction) {
	//store which way the ship is faced
	face := "E"

	var compassL = map[string][]string{
		"E": {"N", "W", "S"},
		"N": {"W", "S", "E"},
		"W": {"S", "E", "N"},
		"S": {"E", "N", "W"},
	}

	var compassR = map[string][]string{
		"E": {"S", "W", "N"},
		"N": {"E", "S", "W"},
		"W": {"N", "E", "S"},
		"S": {"W", "N", "E"},
	}

	//total distance
	var longitude int
	var lattitude int
	direction := face

	for _, s := range directions {

		println("current direction: ", direction)

		if s.direction == "L" {

			steps := s.steps/90 - 1
			possibleDirections := compassL[direction]
			face = possibleDirections[steps]

		} else if s.direction == "R" {

			steps := s.steps/90 - 1
			possibleDirections := compassR[direction]
			face = possibleDirections[steps]

		} else {
			if s.direction != "F" {
				direction = s.direction
			} else {
				direction = face
			}
			switch direction {
			case "N":
				longitude += s.steps
			case "S":
				longitude -= s.steps
			case "W":
				lattitude += s.steps
			case "E":
				lattitude -= s.steps
			}
		}
	}
	total := math.Abs(float64(longitude)) + math.Abs(float64(lattitude))
	fmt.Println(total)
}

func main() {

	directions := parseInputFromFile("example.txt")
	printManhattenDirections(directions)

}
