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

	push func(state bool) bool

	text string // decoded text
)

type app struct{}

func newApp() *app {
	return &app{}
}

func (a *app) Update() error {
	if push == nil {
		push = Push(func(states iter.Seq[bool]) {
			for v := range decode(symbols(pulses(states))) {
				text += v
			}
		})
	}

	on := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace)
	push(on)
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
