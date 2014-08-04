Midi
====

A Go library for receiving MIDI events using ALSA asound.
Specifically designed for use on the Raspberry Pi.

# Dependencies

To install the dependencies in raspbian, install these apt packages:

    sudo apt-get install libasound2-dev libasound2

# Usage

To use the library, you need to call the `Listen` function to set up a MIDI input device.
You can see all your MIDI devices by running:

    aconnect -loi

To connect a device to your app, you use an `aconnect` command such as:

    aconnect 20 128

# Testing

To test your setup, you can use the `midi_dump` tool included in this project.

    go install github.com/bklimt/midi/...
    $GOPATH/bin/midi_dump
