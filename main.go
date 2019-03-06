package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type Timer struct {
	Started  time.Time
	Duration time.Duration
	Beep     bool
}

var black = "\u001b[30m"
var red = "\u001b[31m"
var green = "\u001b[32m"
var yellow = "\u001b[33m"
var blue = "\u001b[34m"
var magenta = "\u001b[35m"
var cyan = "\u001b[36m"
var white = "\u001b[37m"

var bgBlack = "\u001b[40m"
var bgRed = "\u001b[41m"
var bgGreen = "\u001b[42m"
var bgYellow = "\u001b[43m"
var bgBlue = "\u001b[44m"
var bgMagenta = "\u001b[45m"
var bgCyan = "\u001b[46m"
var bgWhite = "\u001b[47m"

var bold = "\u001b[1m"
var reset = "\u001b[0m"
var clean = "\u001b[K"

var spinner = []string{"◜", "◝", "◞", "◟"}
var spinnerIndex = 0

var rendered = false

func main() {
	t := Timer{}

	// parse args
	if len(os.Args) > 1 {
		// parse duration
		d, err := time.ParseDuration(os.Args[1])
		if err == nil {
			t.Duration = d
		}
	}

	t.Reset()
	t.Start()
}

func (t *Timer) Reset() {
	t.Started = time.Now()
	t.Beep = false
}

func (t *Timer) Start() {
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			RenderSpinner()
		}
	}()
	for {
		t.Render()
		time.Sleep(1 * time.Second)
	}
}

func (t *Timer) Render() {

	var h, m, s int64
	numColor := cyan
	bgColor := ""
	target := ""

	if t.Duration == 0 {
		diffTime := time.Since(t.Started)
		h = int64(diffTime.Hours())
		m = int64(diffTime.Minutes()) % 60
		s = int64(diffTime.Seconds()) % 60
	} else {
		x := t.Started.Add(t.Duration)
		diffTime := time.Since(x)
		diffAbs := time.Duration(math.Abs(float64(diffTime)))
		if diffTime > 0 {
			numColor = black
			bgColor = bgRed
			if !t.Beep {
				fmt.Print("\a")
				t.Beep = true
			}
		}

		h = int64(diffAbs.Hours())
		m = int64(diffAbs.Minutes()) % 60
		s = int64(diffAbs.Seconds()) % 60

		target = fmt.Sprintf("%02d:%02d:%02d",
			int64(t.Duration.Hours()),
			int64(t.Duration.Minutes())%60,
			int64(t.Duration.Seconds())%60)
	}

	// reset cursor
	if rendered {
		fmt.Printf("\u001b[9A")
	} else {
		rendered = true
	}

	RenderBorder()

	RenderTargetTime(target)

	// render time
	// up 3
	fmt.Print("\u001b[5A", "\u001b[18G")
	// render number
	fmt.Print(numColor, bgColor, fmt.Sprintf("%02d:%02d:%02d", h, m, s), reset)
	// down 3
	fmt.Print("\u001b[5B", "\u001b[1G")

}

func RenderTargetTime(target string) {
	// up
	fmt.Print("\u001b[2A")
	// render number
	fmt.Print("\u001b[18G", target)
	// down
	fmt.Print("\u001b[2B", "\u001b[1G")
}

func RenderBorder() {
	nl := "\n"
	fmt.Print(clean, "            ╭────────────────╮", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            │                │", nl)
	fmt.Print(clean, "            ╰────────────────╯", nl)
}

func RenderSpinner() {
	// up
	fmt.Print("\u001b[5A")
	// render number
	fmt.Print("\u001b[16G", spinner[spinnerIndex])
	// down
	fmt.Print("\u001b[5B", "\u001b[1G")

	spinnerIndex++
	if spinnerIndex >= len(spinner) {
		spinnerIndex = 0
	}
}
