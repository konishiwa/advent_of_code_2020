package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	startingTime, busses := parseInputFromFile("input.txt")

	getTime(startingTime, busses)

}

func getTime(startingTime string, busses []string) {

	//int startingTime
	time, _ := strconv.Atoi(startingTime)
	//int busses
	busTimes := make([]int, len(busses))
	for i, s := range busses {
		busTimes[i], _ = strconv.Atoi(s)
	}
	getClosestBusses(time, busTimes)

}

func getClosestBusses(startingTime int, busTimes []int) {

	println("startingTime: ", startingTime)
	fmt.Print("busses: ", busTimes)

	//get the modlus of each to get the total waiting time
	waitingTimes := make([]int, len(busTimes))
	//best bus always starts with first
	var bestBusIndex int
	for i, busTime := range busTimes {

		if startingTime%busTime != 0 {
			busStopTime := (startingTime/busTime + 1) * busTime
			println("busStopTime: ", busStopTime)
			waitingTimes[i] = busStopTime - startingTime
		} else {
			println("best bus is: ", busTime)
		}

		if waitingTimes[bestBusIndex] > waitingTimes[i] {
			bestBusIndex = i
		}
	}

	fmt.Print("waitingTime: ", waitingTimes)
	fmt.Print("waitingTime * bus = ", waitingTimes[bestBusIndex]*busTimes[bestBusIndex])

}

func parseInputFromFile(path string) (string, []string) {

	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[0], parseBus(lines[1])
}

func parseBus(input string) []string {

	result := strings.ReplaceAll(input, ",x", "")
	return strings.Split(result, ",")

}
