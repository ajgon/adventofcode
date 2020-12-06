package helpers

type DayAnswers struct {
	Simple   string
	Advanced string
}

type AllAnswers struct {
	Day01 DayAnswers
	Day02 DayAnswers
	Day03 DayAnswers
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
}
