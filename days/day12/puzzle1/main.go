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

func parseInputFromFile() []Direction {

	file, _ := os.Open("input.txt")
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
	//assuming L and R are always 90
	var faceMap = map[string]string{
		"ER90":  "S",
		"EL90":  "N",
		"SL90":  "E",
		"SR90":  "W",
		"NL90":  "W",
		"NR90":  "E",
		"WL90":  "S",
		"WR90":  "N",
		"ER180": "W",
		"EL180": "W",
		"SL180": "N",
		"SR180": "N",
		"NL180": "S",
		"NR180": "S",
		"WL180": "E",
		"WR180": "E",
	}

	//total distance
	var longitude int
	var lattitude int

	for _, s := range directions {

		direction := face

		if s.direction == "L" || s.direction == "R" {
			//println("old direction: ", direction)
			steps := strconv.FormatInt(int64(s.steps), 10)
			key := face + s.direction + steps
			//println("key: ", key)
			//println("s.steps: ", s.steps)
			direction = faceMap[key]
			//println("new direction: ", direction)
		} else {
			if s.direction != "F" {
				direction = s.direction
			}
			println("direction: ", direction)
			println("steps: ", s.steps)
			println("old longitude: ", longitude)
			println("old lattitude: ", lattitude)
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
			println("new longitude: ", longitude)
			println("new lattitude: ", lattitude)
		}
	}
	total := math.Abs(float64(longitude)) + math.Abs(float64(lattitude))
	fmt.Println(total)
}

func main() {

	directions := parseInputFromFile()
	printManhattenDirections(directions)

}
