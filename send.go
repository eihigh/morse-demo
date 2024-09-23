package main

import (
	"fmt"
	"iter"
)

func Send[In, Out any](receive func(iter.Seq[In]) iter.Seq[Out]) (send func(In) iter.Seq[Out], stop func()) {
	var in In
	var consume func(Out) bool

	coro := func(yield func(struct{}) bool) {
		produce := func(pushToRecv func(In) bool) {
			for pushToRecv(in) {
				if !yield(struct{}{}) {
					return
				}
			}
		}
		for result := range receive(produce) {
			if consume == nil {
				return
			}
			if !consume(result) {
				return
			}
		}
	}
	resume, stop := iter.Pull(coro)

	send = func(v In) iter.Seq[Out] {
		return func(yield func(Out) bool) {
			in = v
			consume = yield
			defer func() {
				consume = nil
			}()
			resume() // only one iteration
		}
	}
	return send, stop
}

func sendSample() {
	seq := func(seq iter.Seq[int]) iter.Seq[int] {
		return func(yield func(int) bool) {
			sum := 0
			for v := range seq {
				sum += v
				if v%2 != 0 {
					continue
				}
				if !yield(sum) {
					return
				}
			}
		}
	}

	send, stop := Send(seq)
	defer stop()
	for v := range send(3) {
		fmt.Println("will not print", v)
	}
	for v := range send(4) {
		fmt.Println("will print", v) // v == 7
	}
}
