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
	Day13 DayAnswers
	Day14 DayAnswers
	Day15 DayAnswers
	Day16 DayAnswers
	Day17 DayAnswers
	Day18 DayAnswers
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
	Day13: DayAnswers{
		Simple:   "2382",
		Advanced: "906332393333683",
	},
	Day14: DayAnswers{
		Simple:   "12408060320841",
		Advanced: "4466434626828",
	},
	Day15: DayAnswers{
		Simple:   "203",
		Advanced: "9007186",
	},
	Day16: DayAnswers{
		Simple:   "20231",
		Advanced: "1940065747861",
	},
	Day17: DayAnswers{
		Simple:   "271",
		Advanced: "2064",
	},
	Day18: DayAnswers{
		Simple:   "12918250417632",
		Advanced: "171259538712010",
	},
}
