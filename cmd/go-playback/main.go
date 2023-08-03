package main

import (
	"fmt"
	"os"
	"os/exec"

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
		_tts.TTS(_tts.TEXT, flags.Text, flags.WriteMedia, flags.Voice, flags.Rate, flags.Volume, flags.Proxy, runMPV)
	case flags.File != "":
		_tts.TTS(_tts.FILE, flags.File, flags.WriteMedia, flags.Voice, flags.Rate, flags.Volume, flags.Proxy, runMPV)
	case flags.ListVoices:
		tts.PrintVoices(flags.Proxy)
	default:
		flags.NoExecutableCommand()
	}
}

func runMPV(filename string) {
	cmd := exec.Command("mpv", filename)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
