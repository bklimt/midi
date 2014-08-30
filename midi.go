// Package midi provides a simple API for listening to a subset of events from a MIDI controller.
// It requires ALSA and its development libraries. This package is especially useful on a
// Raspberry Pi.
package midi

// #cgo LDFLAGS: -lasound -lstdc++
// #include "midi.h"
import "C"

// Controller represents a MIDI controller event.
type Controller struct {
	Param int
	Value int
}

// NoteOn represents a MIDI "note on" event.
type NoteOn struct {
	Note int
}

// NoteOff represents a MIDI "note off" event.
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

// Listen starts a goroutine that listens for MIDI events and sends them to the given channel.
func Listen(c chan interface{}) {
	midiChannel = c
	go C.listenForMidiEvents()
}
