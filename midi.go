package midi

// #cgo LDFLAGS: -lasound -lstdc++
// #include "midi.h"
import "C"

type Controller struct {
	Param int
	Value int
}

type NoteOn struct {
	Note int
}

type NoteOff struct {
	Note int
}

var midiChannel chan interface{}

//export onMidiController
func onMidiController(param, value int) {
	midiChannel <- Controller{param, value}
}

//export onMidiNoteOn
func onMidiNoteOn(note int) {
	midiChannel <- NoteOn{note}
}

//export onMidiNoteOff
func onMidiNoteOff(note int) {
	midiChannel <- NoteOff{note}
}

func Listen(c chan interface{}) {
	midiChannel = c
	go C.listenForMidiEvents()
}
