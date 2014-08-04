
#include <stdio.h>
#include <stdlib.h>
#include <alsa/asoundlib.h>
#include <iostream>

extern "C" {
#include "_cgo_export.h"
#include "midi.h"
}

int listenForMidiEvents() {
  std::cout << "Opening default sequencer..." << std::endl;
  snd_seq_t *seq_handle;
  int err = snd_seq_open(&seq_handle, "default", SND_SEQ_OPEN_INPUT, 0);
  if (err) {
    std::cerr << "Unable to open sequencer." << std::endl;
    return -1;
  }

  std::cout << "Setting client name..." << std::endl;
  err = snd_seq_set_client_name(seq_handle, "github.com/bklimt/midi");
  if (err) {
    std::cerr << "Unable to set client name." << std::endl;
    return -2;
  }
  
  std::cout << "Creating port..." << std::endl;
  int in_port = snd_seq_create_simple_port(seq_handle, "listen:in",
      SND_SEQ_PORT_CAP_WRITE | SND_SEQ_PORT_CAP_SUBS_WRITE,
      SND_SEQ_PORT_TYPE_APPLICATION);
  if (in_port < 0) {
    std::cerr << "Invalid port: " << in_port << std::endl;
    return -3;
  }

  std::cout << "Listening for midi events..." << std::endl;

  while (1) {
    snd_seq_event_t *event = NULL;
    snd_seq_event_input(seq_handle, &event);

    std::cout << "Event type: " << (int)event->type << std::endl;
    if (event->type == SND_SEQ_EVENT_CONTROLLER) {
      onMidiController((int)event->data.control.param,
                       (int)event->data.control.value);
    } else if (event->type == SND_SEQ_EVENT_NOTEON) {
      onMidiNoteOn((int)event->data.note.note);
    } else if (event->type == SND_SEQ_EVENT_NOTEOFF) {
      onMidiNoteOff((int)event->data.note.note);
    }
  }

  return 0;
}

