package main

import "iter"

func Push[In any](recv func(iter.Seq[In])) (push func(In) bool) {
	var in In

	coro := func(yieldCoro func(bool) bool) {
		seq := func(yieldSeq func(In) bool) {
			for yieldCoro(yieldSeq(in)) {
			}
		}
		recv(seq)
	}

	next, stop := iter.Pull(coro)
	return func(v In) bool {
		in = v
		more, _ := next()
		if !more {
			stop()
			return false
		}
		return true
	}
}
