package main

import (
	"fmt"
	"github.com/bklimt/midi"
)

func main() {
	c := make(chan interface{})
	midi.Listen(c)
	for event := range c {
		switch event := event.(type) {
		case midi.Controller:
			fmt.Printf("Controller event: %d %d\n", event.Param, event.Value)
		case midi.NoteOn:
			fmt.Printf("Note on: %d\n", event.Note)
		case midi.NoteOff:
			fmt.Printf("Note off: %d\n", event.Note)
		}
	}
}
