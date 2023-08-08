package main

import (
	"github.com/carloseabade/go-tts/internal/flags"
	_tts "github.com/carloseabade/go-tts/internal/tts"
	"github.com/carloseabade/go-tts/pkg/tts"
)

func main() {
	flags.Parse()

	switch {
	case len(flags.Args()) > 0:
		flags.NotACommand()
	case flags.Help:
		flags.Usage()
	case flags.Text != "":
		_tts.TTS(_tts.TEXT, flags.Text, flags.WriteMedia, flags.Voice, flags.Rate, flags.Volume, flags.Proxy)
	case flags.File != "":
		_tts.TTS(_tts.FILE, flags.File, flags.WriteMedia, flags.Voice, flags.Rate, flags.Volume, flags.Proxy)
	case flags.ListVoices:
		tts.PrintVoices(flags.Proxy)
	default:
		flags.NoExecutableCommand()
	}
}
