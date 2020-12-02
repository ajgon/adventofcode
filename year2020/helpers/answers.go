package helpers

type DayAnswers struct {
	Simple   string
	Advanced string
}

type AllAnswers struct {
	Day01 DayAnswers
}

var Answers AllAnswers = AllAnswers{
	Day01: DayAnswers{
		Simple:   "927684",
		Advanced: "292093004",
	},
}
