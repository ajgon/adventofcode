package day04

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day04.input", "\n\n")

type PassportValidator struct {
	Byr *regexp.Regexp
	Iyr *regexp.Regexp
	Eyr *regexp.Regexp
	Hgt *regexp.Regexp
	Hcl *regexp.Regexp
	Ecl *regexp.Regexp
	Pid *regexp.Regexp
}

var passportValidator PassportValidator = PassportValidator{
	Byr: regexp.MustCompile("byr:((19[2-9][0-9])|(200[0-2]))( |$)"),
	Iyr: regexp.MustCompile("iyr:((201[0-9])|2020)( |$)"),
	Eyr: regexp.MustCompile("eyr:((202[0-9])|2030)( |$)"),
	Hgt: regexp.MustCompile("hgt:((((1[5-8][0-9])|(19[0-3]))cm)|((59|(6[0-9])|(7[0-6]))in))( |$)"),
	Hcl: regexp.MustCompile("hcl:#[0-9a-f]{6}( |$)"),
	Ecl: regexp.MustCompile("ecl:(amb|blu|brn|gry|grn|hzl|oth)( |$)"),
	Pid: regexp.MustCompile("pid:[0-9]{9}( |$)"),
}

func simpleValidatePassport(data string) bool {
	data = strings.Join(strings.Split(data, "\n"), " ")
	valid := true

	v := reflect.ValueOf(passportValidator)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		valid = valid && (strings.Index(data, fmt.Sprintf("%s:", strings.ToLower(t.Field(i).Name))) != -1)
	}

	return valid
}

func advancedValidatePassport(data string) bool {
	data = strings.Join(strings.Split(data, "\n"), " ")
	valid := true

	v := reflect.ValueOf(passportValidator)

	for i := 0; i < v.NumField(); i++ {
		re := v.Field(i).Interface().(*regexp.Regexp)
		valid = valid && re.MatchString(data)
	}

	return valid
}

func SimpleSolution() string {
	count := 0

	for _, data := range input {
		if simpleValidatePassport(data) {
			count += 1
		}
	}

	return strconv.Itoa(count)
}

func AdvancedSolution() string {
	count := 0

	for _, data := range input {
		if advancedValidatePassport(data) {
			count += 1
		}
	}

	return strconv.Itoa(count)
}
