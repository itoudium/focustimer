package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Timer struct {
	Started  time.Time
	Border   string
	Duration time.Duration
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

var rendered = false

var width = 30

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
	go func() {
		t.InitBorder(width)
		t.Start()
	}()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		switch text {
		case " ":
			// toggle pause

		case "r", "R":
			// reset timer
		}
	}

}

func (t *Timer) Reset() {
	t.Started = time.Now()
}

func (t *Timer) Start() {
	for {
		t.Render()
		time.Sleep(1 * time.Second)
	}
}

func (t *Timer) InitBorder(w int) {
	t.Border = ""
	for i := 0; i < w; i++ {
		t.Border += "━"
	}
}

func (t *Timer) Render() {

	var h, m, s int64
	numColor := cyan
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
			numColor = red
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
		fmt.Printf("\u001b[7A")
	} else {
		rendered = true
	}

	fmt.Printf(clean+"\r%s\n", t.Border)
	fmt.Printf(clean + "  " + target + "\n")
	fmt.Printf(clean + "\n")
	fmt.Printf(clean+numColor+"           %02d:%02d:%02d\n"+reset, h, m, s)
	fmt.Printf(clean + "\n")
	fmt.Printf(clean + "\n")
	fmt.Printf(clean+"%s\n"+clean, t.Border)

	right := "\u001b[" + strconv.Itoa(width) + "G"
	left := "\u001b[1G"
	up := "\u001b[1A" + left
	down := "\u001b[1B" + right
	v := "┃"
	fmt.Printf(up + "┗")
	for i := 0; i <= 4; i++ {
		fmt.Printf(up + v)
	}
	fmt.Printf(up + "┏" + right + "┓" + down)
	for i := 0; i <= 4; i++ {
		fmt.Printf(v + down)
	}
	fmt.Printf("┛" + down + left)

}