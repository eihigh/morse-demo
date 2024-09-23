# morse-demo
Decoding morse codes interactively with Go's iterators.

![Morse-code-demo](https://github.com/user-attachments/assets/41285cbe-eecf-4ddf-87d5-652bb781cc96)

## How it works
The following code converts a stream of ON/OFF states into text decoded as Morse code.

```go
func newSender() (func(bool) iter.Seq[string], func()) {
	recv := func(states iter.Seq[bool]) iter.Seq[string] {
		return decode(symbols(pulses(states)))
	}
	return Send(recv)
}
```

```go
on := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace)
for s := range send(on) {
    text += s
}
```

TODO: more details
