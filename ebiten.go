package main

import (
	"fmt"
	"iter"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// Dump internal states to visualize the Morse code
	currPulse  pulse
	currSymbol symbol
	currRun    string

	pub Publisher[bool, string] // convert samples to runes

	text string // decoded text
)

type app struct{}

func newApp() *app {
	return &app{}
}

func (a *app) Update() error {
	if pub == nil {
		// Initialize publisher
		sub := func(samples iter.Seq[bool]) iter.Seq[string] {
			return decode(symbols(pulses(samples)))
		}
		pub = Pubsub(sub)
	}

	var on bool
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		on = true
	} else {
		on = false
	}
	out, _ := pub(args2seq(on))
	for _, s := range out {
		text += string(s)
	}
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("Click or Press space\nPulse: %+v\nSymbol: %q\nRun: %s\n%s", currPulse, currSymbol, currRun, text)
	ebitenutil.DebugPrint(screen, msg)
}

func (a *app) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func args2seq[T any](vs ...T) iter.Seq[T] {
	return slices.Values(vs)
}

func main() {
	// Initialize game
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Morse code demo")
	a := newApp()
	if err := ebiten.RunGame(a); err != nil {
		panic(err)
	}
}
