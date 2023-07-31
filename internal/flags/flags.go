package flags

import (
	"flag"
	"fmt"

	"github.com/carloseabade/go-tts/internal/config"
)

var (
	fs             *flag.FlagSet
	Help           bool
	Text           string
	File           string
	Voice          string
	ListVoices     bool
	Rate           string
	Volume         string
	WordsInCue     string
	WriteMedia     string
	WriteSubtitles string
	Proxy          string
)

func init() {
	flag.BoolVar(&Help, "help", false, "show this help message and exit")
	flag.BoolVar(&Help, "h", false, "show this help message and exit")
	flag.StringVar(&Text, "text", "", "what `TEXT` tts will say")
	flag.StringVar(&Text, "t", "", "what `TEXT` tts will say")
	flag.StringVar(&File, "file", "", "tts read text from `FILE`")
	flag.StringVar(&File, "f", "", "tts read text from `FILE`")
	flag.StringVar(&Voice, "voice", "en-US-AriaNeural", "`VOICE` for tts")
	flag.StringVar(&Voice, "v", "en-US-AriaNeural", "`VOICE` for tts")
	flag.BoolVar(&ListVoices, "list-voices", false, "lists available voices and exits")
	flag.BoolVar(&ListVoices, "l", false, "lists available voices and exits")
	flag.StringVar(&Rate, "rate", "+0%", "set tts `RATE`")
	flag.StringVar(&Volume, "volume", "+0%", "set tts `VOLUME`")
	flag.StringVar(&WordsInCue, "words-in-cue", "10", "number of `WORDS` in a subtitle cue")
	flag.StringVar(&WriteMedia, "write-media", "", "send media output to `FILE` instead of stdout")
	flag.StringVar(&WriteSubtitles, "write-subtitles", "", "send subtitle output to provided `FILE` instead of stderr")
	flag.StringVar(&Proxy, "proxy", "", "use a `PROXY` for tts and voice list")
}

func Parse() {
	flag.Parse()
}

func Args() []string {
	return flag.Args()
}

func Usage() {
	fmt.Printf("Usage: %[1]s [OPTIONS]\n\n", config.SoftwareName)
	flag.PrintDefaults()
}

func NotACommand() {
	fmt.Printf("%[1]s: '%[2]s' is not a %[1]s command.\n\n", config.SoftwareName, Args()[0])
	Usage()
}

func NoExecutableCommand() {
	fmt.Printf("%s: error: one of the arguments -t/-text -f/-file -l/-list-voices is required.\n\n", config.SoftwareName)
	Usage()
}
