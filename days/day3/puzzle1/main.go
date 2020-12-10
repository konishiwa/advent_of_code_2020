package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func checkResult(nums []int, num int) {

	for x := 0; x < len(nums); x++ {
		if num < nums[x] {
			//done searching
			break
		} else if num == nums[x] {
			println("The numbers are ", num, " and ", 2020-num)
			println("The answer is: ", num*(2020-num))
			os.Exit(3)
		}
	}
}

func main() {

	//Read the file in as different objects
	inputs, err := ioutil.ReadFile("input.txt")
	check(err)
	tf := string(inputs)
	terrain, err := ReadStrings(strings.NewReader(tf))
	//get the total amount of lines
	terrain_total := len(terrain)
	//get total amount accross
	terrain_across := len(terrain[0])

	//need to loop through to see which ones are needed

	println("terrain_total ", terrain_total)
	println("terrain_across ", terrain_across)
	var total int
	//starting point
	startPoint := terrain[0][0]
	println(string(startPoint))
	total = getTrees(startPoint)
	//slope
	slope := 3
	//need to loop 323 times, always go three over
	for x := 1; x < terrain_total; x++ {
		if slope >= terrain_across {
			slope = slope % terrain_across
		}
		nextTree := terrain[x][slope]
		total += getTrees(nextTree)
		slope = slope + 3
	}

	println("trees ", total)
}

func getTrees(terrains byte) int {

	tree := string(terrains)
	if tree == "#" {
		return 1
	}
	//return 0 if tree = "."
	return 0

}
