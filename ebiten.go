package main

import (
	"fmt"
	"iter"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// Dump internal states to visualize the Morse code
	currDur    int
	currPulse  pulse
	currSymbol symbol
	currRun    string

	push PushFunc[bool, string]

	text string // decoded text
)

type app struct{}

func newApp() *app {
	return &app{}
}

func (a *app) Update() error {
	if push == nil {
		seq := func(states iter.Seq[bool]) iter.Seq[string] {
			return decode(symbols(pulses(states)))
		}
		push = Push(seq)
	}

	on := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace)
	out, _ := push(on, nil)
	for _, s := range out {
		text += s
	}
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("Click or Press space\nUnit: %d\nPulse: %+v\nSymbol: %q\nRun: %s\n%s", currDur/threshold, currPulse, currSymbol, currRun, text)
	ebitenutil.DebugPrint(screen, msg)
}

func (a *app) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 240, 180
}

func main() {
	// Initialize game
	ebiten.SetWindowSize(480, 360)
	ebiten.SetWindowTitle("Morse code demo")
	a := newApp()
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}
}
