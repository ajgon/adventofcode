package main

import (
	"fmt"
	"time"
	"year2020/day01"
	"year2020/day02"
	"year2020/day03"
	"year2020/day04"
	"year2020/day05"
	"year2020/day06"
	"year2020/day07"
	"year2020/day08"
	"year2020/day09"
	"year2020/day10"
	"year2020/day11"
	"year2020/day12"
	"year2020/day13"
	"year2020/day14"

	"github.com/gookit/color"
)

var allAnswersDuration int64 = 0

const solvedDays = 14

func printAnswer(day, variant string, answer string, answerDuration time.Duration) {
	var colonText string
	textColor := color.FgBlue.Render
	variantColor := color.FgYellow.Render
	dayColor := color.FgRed.Render
	answerColor := color.FgGreen.Render
	timeColor := color.FgMagenta.Render

	if variant == "simple" {
		colonText = ":   "
	} else {
		colonText = ": "
	}

	color.Printf(
		"%s%s%s%s%s%s%s\n",
		textColor("Day "),
		dayColor(day),
		textColor(" "),
		variantColor("("+variant+")"),
		textColor(colonText),
		answerColor(fmt.Sprintf("%s", answer)),
		timeColor(fmt.Sprintf(" (%v)", answerDuration)),
	)
}

func handleAnswer(answerCall func() string, answerChan chan string, durationChan chan time.Duration) {
	start := time.Now()
	answer, duration := answerCall(), time.Since(start)
	answerChan <- answer
	durationChan <- duration
}

var daysMap = []func() string{
	day01.SimpleSolution, day01.AdvancedSolution,
	day02.SimpleSolution, day02.AdvancedSolution,
	day03.SimpleSolution, day03.AdvancedSolution,
	day04.SimpleSolution, day04.AdvancedSolution,
	day05.SimpleSolution, day05.AdvancedSolution,
	day06.SimpleSolution, day06.AdvancedSolution,
	day07.SimpleSolution, day07.AdvancedSolution,
	day08.SimpleSolution, day08.AdvancedSolution,
	day09.SimpleSolution, day09.AdvancedSolution,
	day10.SimpleSolution, day10.AdvancedSolution,
	day11.SimpleSolution, day11.AdvancedSolution,
	day12.SimpleSolution, day12.AdvancedSolution,
	day13.SimpleSolution, day13.AdvancedSolution,
	day14.SimpleSolution, day14.AdvancedSolution,
}

func main() {
	start := time.Now()
	dayChans := make([]chan string, 0)
	durationChans := make([]chan time.Duration, 0)
	summarizedDuration := 0

	for i := 0; i < len(daysMap); i++ {
		newDayChan := make(chan string)
		dayChans = append(dayChans, newDayChan)

		newDurationChan := make(chan time.Duration)
		durationChans = append(durationChans, newDurationChan)

		go handleAnswer(daysMap[i], newDayChan, newDurationChan)
	}

	color.Style{color.FgMagenta, color.OpBold}.Println("Year 2020 answers\n")

	for i := 0; i < len(daysMap); i++ {
		variant := ""
		if i%2 == 0 {
			variant = "simple"
		} else {
			variant = "advanced"
		}
		solution := <-dayChans[i]
		duration := <-durationChans[i]
		printAnswer(fmt.Sprintf("%02d", i/2+1), variant, solution, duration)

		summarizedDuration += int(duration.Nanoseconds())
	}

	color.Style{color.FgMagenta, color.OpBold}.Println(fmt.Sprintf("\nSummarized duration of all days: %v", time.Duration(summarizedDuration)))
	color.Style{color.FgMagenta, color.OpBold}.Println(fmt.Sprintf("Parallel processing time: %v", time.Since(start)))
}
