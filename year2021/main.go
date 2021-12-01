package main

import (
	"fmt"
	"time"

	"github.com/gookit/color"
	"year2021/day01"
)

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
		answerColor(answer),
		timeColor(fmt.Sprintf(" (%v)", answerDuration)),
	)
}

func handleAnswer(answerCall func() string, answerChan chan string, durationChan chan time.Duration) {
	start := time.Now()
	answer, duration := answerCall(), time.Since(start)
	answerChan <- answer
	durationChan <- duration
}

func main() {
	daysMap := []func() string{
		day01.SimpleSolution, day01.AdvancedSolution,
	}

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

	color.Style{color.FgMagenta, color.OpBold}.Println("Year 2021 answers\n")

	for i := 0; i < len(daysMap); i++ {
		var variant string
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

	color.Style{color.FgMagenta, color.OpBold}.Println(
		fmt.Sprintf("\nSummarized duration of all days: %v", time.Duration(summarizedDuration)),
	)
	color.Style{color.FgMagenta, color.OpBold}.Println(
		fmt.Sprintf("Parallel processing time: %v", time.Since(start)),
	)
}
