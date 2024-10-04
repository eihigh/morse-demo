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

	text   string // decoded text
	textJP string
)

type app struct{}

func newApp() *app {
	return &app{}
}

func (a *app) Update() error {
	if push == nil {
		jp := Push(func(symbols iter.Seq[symbol]) {
			for v := range decodeJP(symbols) {
				textJP += v
				fmt.Println(textJP)
			}
		})

		en := Push(func(symbols iter.Seq[symbol]) {
			for v := range decode(symbols) {
				text += v
			}
		})

		push = Push(func(states iter.Seq[bool]) {
			for v := range symbols(pulses(states)) {
				if !jp(v) || !en(v) {
					return
				}
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
