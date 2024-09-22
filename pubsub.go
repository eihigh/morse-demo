package main

import (
	"iter"
)

type Subscriber[VIn, VOut any] func(iter.Seq[VIn]) iter.Seq[VOut]

type Publisher[VIn, VOut any] func(iter.Seq[VIn]) ([]VOut, bool)

func Pubsub[VIn, VOut any](subscriber Subscriber[VIn, VOut]) Publisher[VIn, VOut] {
	var in VIn
	var out []VOut

	coro := func(yieldCoro func(more bool) bool) {
		seq := subscriber(func(yieldSeq func(VIn) bool) {
			for {
				more := yieldSeq(in)
				if !yieldCoro(more) {
					break
				}
			}
		})

		for v := range seq {
			out = append(out, v)
		}
	}

	next, stop := iter.Pull(coro)
	return func(incoming iter.Seq[VIn]) ([]VOut, bool) {
		out = nil
		var more bool
		for v := range incoming {
			in = v
			more, _ = next()
			if !more {
				break
			}
		}
		if more {
			return out, true
		}
		stop()
		return out, false
	}
}
