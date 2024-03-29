package helpers

type DayAnswers struct {
	Simple   string
	Advanced string
}

type AllAnswers struct {
	Day01 DayAnswers
	Day02 DayAnswers
}

func Answers() AllAnswers {
	return AllAnswers{
		Day01: DayAnswers{
			Simple:   "1477",
			Advanced: "1523",
		},
		Day02: DayAnswers{
			Simple:   "1670340",
			Advanced: "1954293920",
		},
	}
}
