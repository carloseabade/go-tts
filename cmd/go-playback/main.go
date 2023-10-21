package main

import (
	"fmt"
	"os/exec"

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
		_tts.TTSPlayback(_tts.TEXT, flags.GetText(), flags.GetWriteMedia(), flags.GetVoice(), flags.GetRate(), flags.GetVolume(), flags.GetProxy(), flags.GetFast(), runMPV)
	case flags.GetFile() != "":
		_tts.TTSPlayback(_tts.FILE, flags.GetFile(), flags.GetWriteMedia(), flags.GetVoice(), flags.GetRate(), flags.GetVolume(), flags.GetProxy(), flags.GetFast(), runMPV)
	case flags.GetListVoices():
		tts.PrintVoices(flags.GetProxy())
	default:
		flags.NoExecutableCommand()
	}
}

func runMPV(filename string) {
	cmd := exec.Command("mpv", filename)

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
	}
}
