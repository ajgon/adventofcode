package helpers

type DayAnswers struct {
	Simple   string
	Advanced string
}

type AllAnswers struct {
	Day01 DayAnswers
	Day02 DayAnswers
	Day03 DayAnswers
	Day04 DayAnswers
	Day05 DayAnswers
	Day06 DayAnswers
	Day07 DayAnswers
	Day08 DayAnswers
	Day09 DayAnswers
	Day10 DayAnswers
	Day11 DayAnswers
	Day12 DayAnswers
}

var Answers AllAnswers = AllAnswers{
	Day01: DayAnswers{
		Simple:   "927684",
		Advanced: "292093004",
	},
	Day02: DayAnswers{
		Simple:   "416",
		Advanced: "688",
	},
	Day03: DayAnswers{
		Simple:   "220",
		Advanced: "2138320800",
	},
	Day04: DayAnswers{
		Simple:   "242",
		Advanced: "186",
	},
	Day05: DayAnswers{
		Simple:   "885",
		Advanced: "623",
	},
	Day06: DayAnswers{
		Simple:   "6799",
		Advanced: "3354",
	},
	Day07: DayAnswers{
		Simple:   "124",
		Advanced: "34862",
	},
	Day08: DayAnswers{
		Simple:   "1394",
		Advanced: "1626",
	},
	Day09: DayAnswers{
		Simple:   "400480901",
		Advanced: "67587168",
	},
	Day10: DayAnswers{
		Simple:   "2176",
		Advanced: "18512297918464",
	},
	Day11: DayAnswers{
		Simple:   "2299",
		Advanced: "2047",
	},
	Day12: DayAnswers{
		Simple:   "1133",
		Advanced: "61053",
	},
}
