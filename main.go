package main

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/ojrac/opensimplex-go"
)

func main() {
	n := opensimplex.New()
	termbox.Init()
	termbox.SetOutputMode(termbox.OutputGrayscale)
	defer termbox.Close()
	width, height := termbox.Size()

	go func() {
		t := 0
		scale := 20.0
		ticker := time.NewTicker(50 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				t++
				for x := 0; x < width; x++ {
					for y := 0; y < height; y++ {
						c := 30 + int(n.Eval3(float64(x)/scale, float64(y)/scale, float64(t)/scale)*20)
						termbox.SetCell(x, y, '$', termbox.Attribute(c), termbox.ColorDefault)
					}
				}
				termbox.Flush()
			}
		}
	}()

	termbox.PollEvent()

}
