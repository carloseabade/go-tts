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

//edge-tts --text "Hello, world!" --write-media hello.mp3 --write-subtitles hello.vtt
//edge-playback --text "Hello, world!"
//edge-tts --list-voices
//edge-tts --voice ar-EG-SalmaNeural --text "مرحبا كيف حالك؟" --write-media hello_in_arabic.mp3 --write-subtitles hello_in_arabic.vtt
//edge-tts --rate=-50% --text "Hello, world!" --write-media hello_with_rate_halved.mp3 --write-subtitles hello_with_rate_halved.vtt
//edge-tts --volume=-50% --text "Hello, world!" --write-media hello_with_volume_halved.mp3 --write-subtitles hello_with_volume_halved.vtt
