# morse-demo
Decoding morse codes interactively with Go's iterators.

![Morse-code-demo](https://github.com/user-attachments/assets/41285cbe-eecf-4ddf-87d5-652bb781cc96)

## How it works
The following code converts a stream of ON/OFF states into text decoded as Morse code.

```go
// newPublisher creates a new Publisher that decodes a stream of boolean states.
func newPublisher() Publisher[bool, string] {
	sub := func(states iter.Seq[bool]) iter.Seq[string] {
		return decode(symbols(pulses(states)))
	}
	return Pubsub(sub)
}

// publishToDecode publishes a state to a Publisher and returns the decoded text.
func publishToDecode(pub Publisher[bool, string], states iter.Seq[bool]) []string {
	out, _ := pub(states)
	return out
}
```

TODO: more details
