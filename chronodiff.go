package main

import (
	"flag"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	var (
		start string
		end   string
		pause int
	)

	flag.StringVar(&start, "start", "07:00", "Startzeit")
	flag.StringVar(&end, "end", "15:00", "Endzeit")
	flag.IntVar(&pause, "break", 45, "Pause in Minuten")

	flag.Parse()

	start = start + ":00"
	end = end + ":00"

	startTime, _ := time.Parse(time.TimeOnly, start)
	endTime, _ := time.Parse(time.TimeOnly, end)

	workDurationNoBreak := endTime.Sub(startTime)
	workDuration := workDurationNoBreak - time.Minute*time.Duration(pause)
	industrialNoBreak := math.Round((workDurationNoBreak.Hours() * 100)) / 100
	industrial := math.Round((workDuration.Hours() * 100)) / 100

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Normal", "Industrial"})
	t.AppendRows([]table.Row{
		{"Without break", workDurationNoBreak, industrialNoBreak},
		{"With break", workDuration, industrial},
	})
	t.Render()
}
