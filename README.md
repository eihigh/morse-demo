# morse-demo
Decoding morse codes interactively with Go's iterators.

![Morse-code-demo](https://github.com/user-attachments/assets/41285cbe-eecf-4ddf-87d5-652bb781cc96)

## How it works
The following code converts a stream of ON/OFF states into text decoded as Morse code.
```go
	if push == nil {
		push = Push(func(states iter.Seq[bool]) {
			for v := range decode(symbols(pulses(states))) {
				text += v
			}
		})
	}

	on := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace)
	push(on)
```
TODO: more details
