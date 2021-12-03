package main

import (
	"image"
	"log"
	"time"
)
import "github.com/go-vgo/robotgo"
import "github.com/vcaesar/gcv"

type Event struct {
	Img         image.Image
	DoubleClick bool
	Text        string
}

func (e *Event) Do() {
	if e.Img != nil {
		sc := robotgo.CaptureScreen()
		defer robotgo.FreeBitmap(sc)
		img := robotgo.ToImage(sc)
		x, y := gcv.FindImgXY(e.Img, img)
		x = x / 2
		y = y / 2
		robotgo.MoveClick(x, y, "left", e.DoubleClick)
	}
	if e.Text != "" {
		time.Sleep(time.Second)
		for _, i := range e.Text {
			e.key_press(string(i))
		}
		e.key_press("enter")
	}
}
func (e *Event) key_press(s string) {
	robotgo.KeyDown(s)
	robotgo.KeyUp(s)
}


//go:generate go get -d github.com/goware/modvendor
//go:generate go mod vendor
//go:generate modvendor -copy="**/*.c **/*.h **/*.proto **/*.a" -v
func main() {
	img, _, err := robotgo.DecodeImg("assets/terminal.png")
	if err != nil {
		log.Fatal(err)
	}

	es := []Event{
		{img, false, ""},
		{nil, false, "ps -ef"},
	}

	for _, e := range es {
		e.Do()
	}
}

