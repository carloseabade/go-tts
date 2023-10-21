package main

import (
	"github.com/carloseabade/go-tts/internal/flags"
	_tts "github.com/carloseabade/go-tts/internal/tts"
	"github.com/carloseabade/go-tts/pkg/tts"
)

func main() {
	flags.Parse()

	switch {
	case flags.HasNonFlagArguments():
		flags.NotACommand()
	case flags.GetHelp():
		flags.Usage()
	case flags.GetText() != "":
		_tts.TTS(_tts.TEXT, flags.GetText(), flags.GetWriteMedia(), flags.GetVoice(), flags.GetRate(), flags.GetVolume(), flags.GetProxy())
	case flags.GetFile() != "":
		_tts.TTS(_tts.FILE, flags.GetFile(), flags.GetWriteMedia(), flags.GetVoice(), flags.GetRate(), flags.GetVolume(), flags.GetProxy())
	case flags.GetListVoices():
		tts.PrintVoices(flags.GetProxy())
	default:
		flags.NoExecutableCommand()
	}
}
