package day19

import (
	"regexp"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day19.input")

func extractRulesAndCandidates(data []string) (rules []string, candidates []string) {
	appendRules := true
	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			appendRules = false
		}

		if appendRules {
			rules = append(rules, data[i])
		} else {
			candidates = append(candidates, data[i])
		}
	}

	return
}

func normalizeRules(ruleset []string) string {
	lettersOnlyRegex := regexp.MustCompile("^([0-9]+): \"?([^0-9\"]+)\"? +$")
	canBreak := true

	for i := 0; i < len(ruleset); i++ {
		ruleset[i] = ruleset[i] + " "
	}

	for {
		canBreak = true
		for i := 0; i < len(ruleset); i++ {
			matches := lettersOnlyRegex.FindStringSubmatch(ruleset[i])
			if len(matches) > 0 {
				for j := 0; j < len(ruleset); j++ {
					if len(matches[2]) == 1 {
						ruleset[j] = strings.ReplaceAll(strings.ReplaceAll(ruleset[j], " "+matches[1]+" ", " "+matches[2]+" "), " "+matches[1]+" ", " "+matches[2]+" ")

					} else {
						ruleset[j] = strings.ReplaceAll(strings.ReplaceAll(ruleset[j], " "+matches[1]+" ", " (?:"+matches[2]+") "), " "+matches[1]+" ", " (?:"+matches[2]+") ")

					}
				}
			} else {
				canBreak = false
			}
		}

		if canBreak {
			break
		}
	}

	for i := 0; i < len(ruleset); i++ {
		if ruleset[i][0:2] == "0:" {
			return "^" + strings.ReplaceAll(ruleset[i][3:], " ", "") + "$"
		}
	}

	return ""
}

func permutateRules(rules []string, deepness int) [][]string {
	results := make([][]string, 0)

	for r8 := 0; r8 < deepness; r8++ {
		for r11 := 0; r11 < deepness; r11++ {
			for i := 0; i < len(rules); i++ {
				newRule := ""
				if rules[i][0:2] == "8:" {
					for j := 0; j < r8+1; j++ {
						newRule += " 42"
					}
					rules[i] = "8:" + newRule
				}

				if rules[i][0:3] == "11:" {
					for j := 0; j < r11+1; j++ {
						newRule += " 42"
					}
					for j := 0; j < r11+1; j++ {
						newRule += " 31"
					}
					rules[i] = "11:" + newRule
				}
			}
			newRules := make([]string, len(rules))
			copy(newRules, rules)
			results = append(results, newRules)
		}
	}

	return results
}

func findSimpleMatches(data []string, rulesRegex string) []string {
	matches := make([]string, 0)
	re := regexp.MustCompile(rulesRegex)

	for i := 0; i < len(data); i++ {
		if re.MatchString(data[i]) {
			matches = append(matches, data[i])
		}
	}

	return matches
}

func findAdvancedMatches(candidates []string, rules []string) []string {
	permutations := permutateRules(rules, 5)
	finalResults := make([]string, 0)

	for p := 0; p < len(permutations); p++ {
		results := findSimpleMatches(candidates, normalizeRules(permutations[p]))

		for r := 0; r < len(results); r++ {
			if !helpers.ArrayContainsStr(finalResults, results[r]) {
				finalResults = append(finalResults, results[r])
			}
		}
	}

	return finalResults
}

func SimpleSolution() string {
	rules, candidates := extractRulesAndCandidates(input)

	return strconv.Itoa(len(findSimpleMatches(candidates, normalizeRules(rules))))
}

func AdvancedSolution() string {
	rules, candidates := extractRulesAndCandidates(input)

	return strconv.Itoa(len(findAdvancedMatches(candidates, rules)))
}
