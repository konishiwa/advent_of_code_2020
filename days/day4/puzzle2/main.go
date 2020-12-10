package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func hasSymbol(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Read the file in as different objects
	inputs, err := ioutil.ReadFile("input.txt")
	check(err)
	tf := string(inputs)
	temp := strings.Split(tf, "\n\n")
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	results := 0

	for x := 0; x < len(temp); x++ {
		// Should replace all newlines with a space delimted keypairs
		temp[x] = strings.ReplaceAll(temp[x], "\n", " ")
		// println(temp[x])
		valid := true
		// println(temp[x] + "\n")
		//make sure it has all kays
		for i := 0; i < len(keys); i++ {
			if !(strings.Contains(temp[x], keys[i])) {
				valid = false
			}
		}

		attr := strings.Split(temp[x], " ")

		numMap := make(map[string]string)
		//create dict
		for x := 0; x < len(attr); x++ {
			field := strings.Split(attr[x], ":")
			numMap[field[0]] = field[1]
		}
		valid = validator(numMap)
		if valid == true {
			results++
		}
	}
	println("valid passportds are: ", results)
}

func validator(passport map[string]string) bool {

	//validate byr
	if byr, ok := passport["byr"]; ok {
		byrInt, err := strconv.Atoi(byr)
		if err != nil {
			println("bad byr conversion", byr)
			return false
		}
		if byrInt < 1920 || byrInt > 2002 {
			// println("byr int not between 1920-2002: ", byrInt)
			return false
		}
	} else {
		return false
	}

	//validate iyr
	if iyr, ok := passport["iyr"]; ok {
		iyrInt, err := strconv.Atoi(iyr)
		if err != nil {
			println("bad iyr conversion", iyr)
			return false
		}
		if iyrInt < 2010 || iyrInt > 2020 {
			// println("iyr int not between 2010-2020: ", iyrInt)
			return false
		}
	} else {
		return false
	}

	//validate eyr
	if eyr, ok := passport["eyr"]; ok {
		eyrInt, err := strconv.Atoi(eyr)
		if err != nil {
			println("bad eyr conversion", eyr)
			return false
		}
		if eyrInt < 2020 || eyrInt > 2030 {
			// println("eyr int not between 2020-2030: ", eyrInt)
			return false
		}
	} else {
		return false
	}

	//hgt validate
	if hgt, ok := passport["hgt"]; ok {
		// A sketchy call on my part... lots of ways to break it..
		num := hgt[:len(hgt)-2]
		h, err := strconv.Atoi(num)
		if err != nil {
			println("bad hgt conversion", num)
			return false
		}
		// Ensures height is followed/ended with cm/in
		switch string(hgt[len(hgt)-2:]) {
		case "cm":
			if h > 193 || h < 150 {
				println("cm invalid hgt", num)
				return false
			}
		case "in":
			if h > 76 || h < 59 {
				println("in invalid hgt", num)
				return false
			}
		default:
			println("bad hgt", num)
		}
	} else {
		return false
	}
	// if string(hgt[len(hgt)-2:]) == "cm" || string(hgt[len(hgt)-2:]) == "in" {
	// 	num := hgt[:len(hgt)-2]
	// 	hgtInt, err := strconv.Atoi(num)
	// 	if err != nil {
	// 		return false
	// 	}
	// 	if strings.Contains(hgt, "in") {
	// 		println("height contains inches ", hgt)
	// 		if hgtInt > 76 || hgtInt < 59 {
	// 			return false
	// 		}
	// 	} else {
	// 		if hgtInt > 193 || hgtInt < 150 {
	// 			return false
	// 		}
	// 	}
	// } else {
	// 	return false
	// }

	//hcl validate
	if hcl, ok := passport["hcl"]; ok {
		//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		// if 0 != strings.Index(hcl, "#") || len(hcl) != 7 {
		if 0 != strings.Index(hcl, "#") || len(hcl) != 7 {
			println("bad hcl", hcl)
			return false
		}
		if hasSymbol(hcl[:len(hcl)-1]) {
			println("hcl contains invalid input ", hcl)
			return false
		}
	} else {
		return false
	}

	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if ecl, ok := passport["ecl"]; ok {

		//should be a dict
		validEcl := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

		eclfound := false

		for x := 0; x < len(validEcl); x++ {
			if validEcl[x] == ecl[:3] {
				eclfound = true
			}
		}
		if eclfound == false {
			return false
		}

	} else {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if pid, ok := passport["pid"]; ok {
		if len(pid) != 9 {
			// println("invalid pid", pid)
			return false
		}
		_, errPid := strconv.Atoi(pid)
		if errPid != nil {
			println("its not a right pid ", pid)
			return false
		}
	} else {
		return false
	}

	return true
}
