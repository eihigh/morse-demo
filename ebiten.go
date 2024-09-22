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
	currDur    int
	currPulse  pulse
	currSymbol symbol
	currRun    string

	pub Publisher[bool, string] // convert states to runes

	text string // decoded text
)

type app struct{}

func newApp() *app {
	return &app{}
}

func (a *app) Update() error {
	if pub == nil {
		pub = newPublisher()
	}

	on := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace)
	for _, s := range publishToDecode(pub, args2seq(on)) {
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

func args2seq[T any](vs ...T) iter.Seq[T] {
	return slices.Values(vs)
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
