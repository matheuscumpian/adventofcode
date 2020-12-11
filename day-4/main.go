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

var requiredPassportFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entries := readFile(scanner)
	fmt.Println(countValidPassports(entries))
}

func countValidPassports(passports []map[string]string) int {
	var valids int
	for _, passport := range passports {
		if isValidPassport(passport) {
			valids++
		}
	}
	return valids
}

func isValidField(field string, value string) bool {
	switch field {
	case "byr":
		if len(field) != 3 {
			return false
		}
		birth, _ := strconv.ParseInt(value, 10, 64)
		if birth < 1920 || birth > 2002 {
			return false
		}
		break
	case "iyr":
		if len(field) != 3 {
			return false
		}
		issue, _ := strconv.ParseInt(value, 10, 64)
		if issue < 2010 || issue > 2020 {
			return false
		}
		break
	case "eyr":
		if len(field) != 3 {
			return false
		}
		expiration, _ := strconv.ParseInt(value, 10, 64)
		if expiration < 2020 || expiration > 2030 {
			return false
		}
		break
	case "hgt":
		if !strings.Contains(value, "cm") && !strings.Contains(value, "in") {
			return false
		}

		if strings.Contains(value, "cm") {
			heightString := strings.Split(value, "cm")[0]
			height, _ := strconv.ParseInt(heightString, 10, 64)
			if height < 150 || height > 193 {
				return false
			}
		} else if strings.Contains(value, "in") {
			heightString := strings.Split(value, "in")[0]
			height, _ := strconv.ParseInt(heightString, 10, 64)
			if height < 59 || height > 76 {
				return false
			}
		}
	case "hcl":
		r, _ := regexp.Compile("^#([0-9a-f]){6}")
		if !r.MatchString(value) {
			return false
		}
	case "ecl":
		if !isStringOnArray(validEyeColors, value) {
			return false
		}
	case "pid":
		r, _ := regexp.Compile("^\\d{9}$")
		if !r.MatchString(value) {
			return false
		}
	}

	return true
}

func isValidPassport(passport map[string]string) bool {
	for _, field := range requiredPassportFields {
		if _, ok := passport[field]; !ok {
			return false
		}
		if !isValidField(field, passport[field]) {
			return false
		}
	}
	return true
}

func readFile(s *bufio.Scanner) []map[string]string {
	var mapArray []map[string]string
	for s.Scan() {
		m := make(map[string]string)
		for {
			lineText := s.Text()
			if len(lineText) == 0 {
				break
			}
			pairs := strings.Split(lineText, " ")
			for _, pair := range pairs {
				arr := strings.Split(pair, ":")
				key := arr[0]
				value := arr[1]
				m[key] = value
			}
			s.Scan()
		}
		mapArray = append(mapArray, m)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	return mapArray
}

func isStringOnArray(arr []string, value string) bool {
	for _, val := range arr {
		if val == value {
			return true
		}
	}
	return false
}
