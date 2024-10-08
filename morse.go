package main

import (
	"iter"
	"slices"
)

var (
	threshold = 12 // in ticks
)

func args[T any](vs ...T) iter.Seq[T] {
	return slices.Values(vs)
}

type pulse struct {
	on       bool
	duration int
}

func pulses(states iter.Seq[bool]) iter.Seq[pulse] {
	return func(yield func(pulse) bool) {
		t := 0
		prevT := 0
		prevOn := false
		for on := range states {
			currDur = t - prevT // for visualization
			if on != prevOn {
				d := t - prevT
				p := pulse{prevOn, d}
				prevT = t
				prevOn = on
				currPulse = p // for visualization
				if !yield(p) {
					return
				}
			}
			t++
		}
	}
}

type symbol string

const (
	dot    symbol = "."
	dash   symbol = "-"
	letter symbol = "letter"
	space  symbol = "space"
)

func symbols(pulses iter.Seq[pulse]) iter.Seq[symbol] {
	return func(yield func(symbol) bool) {
		for pulse := range pulses {
			var sym symbol
			if pulse.on {
				if pulse.duration < threshold {
					sym = dot
				} else {
					sym = dash
				}
			} else {
				if pulse.duration < threshold*2 {
					continue // discard
				} else if pulse.duration < threshold*6 {
					sym = letter
				} else {
					sym = space
				}
			}
			currSymbol = sym // for visualization
			if !yield(sym) {
				return
			}
		}
	}
}

func decode(syms iter.Seq[symbol]) iter.Seq[string] {
	return func(yield func(string) bool) {
		run := ""
		for sym := range syms {
			if sym == letter || sym == space {
				// flush
				if run != "" {
					ch, ok := table[run]
					if !ok {
						ch = "?"
					}
					if !yield(ch) {
						return
					}
					run = ""
				}
				if sym == space {
					if !yield(" ") {
						return
					}
				}
			} else {
				run += string(sym)
			}
			currRun = run // for visualization
		}
	}
}

func decodeJP(syms iter.Seq[symbol]) iter.Seq[string] {
	return func(yield func(string) bool) {
		run := ""
		for sym := range syms {
			if sym == letter || sym == space {
				// flush
				if run != "" {
					ch, ok := tableJP[run]
					if !ok {
						ch = "?"
					}
					if !yield(ch) {
						return
					}
					run = ""
				}
				if sym == space {
					if !yield(" ") {
						return
					}
				}
			} else {
				run += string(sym)
			}
			currRun = run // for visualization
		}
	}
}

var table = map[string]string{
	".-":   "A",
	"-...": "B",
	"-.-.": "C",
	"-..":  "D",
	".":    "E",
	"..-.": "F",
	"--.":  "G",
	"....": "H",
	"..":   "I",
	".---": "J",
	"-.-":  "K",
	".-..": "L",
	"--":   "M",
	"-.":   "N",
	"---":  "O",
	".--.": "P",
	"--.-": "Q",
	".-.":  "R",
	"...":  "S",
	"-":    "T",
	"..-":  "U",
	"...-": "V",
	".--":  "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",
}

var tableJP = map[string]string{
	".-":     "イ",
	".-.-":   "ロ",
	"-...":   "ハ",
	"-.-.":   "ニ",
	"-..":    "ホ",
	".":      "ヘ",
	"..-..":  "ト",
	"..-.":   "チ",
	"--.":    "リ",
	"....":   "ヌ",
	"-.--.":  "ル",
	".---":   "ヲ",
	"-.-":    "ワ",
	".-..":   "カ",
	"--":     "ヨ",
	"-.":     "タ",
	"---":    "レ",
	"---.":   "ソ",
	".--.":   "ツ",
	"--.-":   "ネ",
	".-.":    "ナ",
	"...":    "ラ",
	"-":      "ム",
	"..-":    "ウ",
	".-..-":  "ヰ",
	"..--":   "ノ",
	".-...":  "オ",
	"...-":   "ク",
	".--":    "ヤ",
	"-..-":   "マ",
	"-.--":   "ケ",
	"--..":   "フ",
	"----":   "コ",
	"-.---":  "エ",
	".-.--":  "テ",
	"--.--":  "ア",
	"-.-.-":  "サ",
	"-.-..":  "キ",
	"-..--":  "ユ",
	"-...-":  "メ",
	"..-.--": "ミ",
	"--.-.":  "シ",
	".--..":  "ヱ",
	"--..-":  "ヒ",
	"-..-.":  "モ",
	".---.":  "セ",
	"---.-":  "ス",
	".-.-.":  "ン",
	"..":     "゛",
	"..--.":  "゜",
	".--.-":  "ー",
}
