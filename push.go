package main

import "iter"

type SeqFunc[VIn, VOut any] func(iter.Seq[VIn]) iter.Seq[VOut]

type PushFunc[VIn, VOut any] func(in VIn, outbuf []VOut) ([]VOut, bool)

func Push[VIn, VOut any](seqfunc SeqFunc[VIn, VOut]) PushFunc[VIn, VOut] {
	var tmpIn VIn
	var tmpOutbuf []VOut

	coro := func(yieldCoro func(bool) bool) {
		seq := seqfunc(func(yieldSeq func(VIn) bool) {
			for yieldCoro(yieldSeq(tmpIn)) {
			}
		})

		for v := range seq {
			tmpOutbuf = append(tmpOutbuf, v)
		}
	}

	next, stop := iter.Pull(coro)
	return func(in VIn, outbuf []VOut) ([]VOut, bool) {
		tmpIn = in
		tmpOutbuf = outbuf
		more, _ := next()
		if !more {
			stop()
			return tmpOutbuf, false
		}
		return tmpOutbuf, true
	}
}
