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
}
