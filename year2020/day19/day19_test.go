package day19

import (
	"reflect"
	"sort"
	"strings"
	"testing"
	"year2020/helpers"
)

var testSimpleRules = strings.Split(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`, "\n")

var testSimpleData = strings.Split(`ababbb
bababa
abbbab
aaabbb
aaaabbb`, "\n")

var testAdvancedRules = strings.Split(`42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1`, "\n")

var testAdvancedData = strings.Split(`abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`, "\n")

func TestNormalizeRules(t *testing.T) {
	got := normalizeRules(testSimpleRules)
	want := "^a(?:(?:aa|bb)(?:ab|ba)|(?:ab|ba)(?:aa|bb))b$"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestFindMatches(t *testing.T) {
	rulesRegex := normalizeRules(testSimpleRules)
	got := findSimpleMatches(testSimpleData, rulesRegex)
	want := []string{"ababbb", "abbbab"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPermutateRules(t *testing.T) {
	permutations := permutateRules(testAdvancedRules, 3)
	got := []string{
		permutations[0][26], permutations[0][4],
		permutations[1][26], permutations[1][4],
		permutations[2][26], permutations[2][4],
		permutations[3][26], permutations[3][4],
		permutations[4][26], permutations[4][4],
		permutations[5][26], permutations[5][4],
		permutations[6][26], permutations[6][4],
		permutations[7][26], permutations[7][4],
		permutations[8][26], permutations[8][4],
	}
	want := []string{
		"8: 42", "11: 42 31",
		"8: 42", "11: 42 42 31 31",
		"8: 42", "11: 42 42 42 31 31 31",
		"8: 42 42", "11: 42 31",
		"8: 42 42", "11: 42 42 31 31",
		"8: 42 42", "11: 42 42 42 31 31 31",
		"8: 42 42 42", "11: 42 31",
		"8: 42 42 42", "11: 42 42 31 31",
		"8: 42 42 42", "11: 42 42 42 31 31 31",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindAdvancedMatches(t *testing.T) {
	got := findAdvancedMatches(testAdvancedData, testAdvancedRules)
	want := []string{
		"bbabbbbaabaabba",
		"babbbbaabbbbbabbbbbbaabaaabaaa",
		"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
		"bbbbbbbaaaabbbbaaabbabaaa",
		"bbbababbbbaaaaaaaabbababaaababaabab",
		"ababaaaaaabaaab",
		"ababaaaaabbbaba",
		"baabbaaaabbaaaababbaababb",
		"abbbbabbbbaaaababbbbbbaaaababb",
		"aaaaabbaabaaaaababaa",
		"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
		"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
	}
	sort.Strings(got)
	sort.Strings(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day19.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day19.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
