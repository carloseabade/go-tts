package flags

import (
	"flag"
	"fmt"

	"github.com/carloseabade/go-tts/internal/config"
)

var (
	fs         *flag.FlagSet
	help       bool
	text       string
	file       string
	voice      string
	listVoices bool
	rate       string
	volume     string
	//wordsInCue     string
	writeMedia string
	//writeSubtitles string
	proxy string
	fast  bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show this help message and exit")
	flag.BoolVar(&help, "h", false, "show this help message and exit")
	flag.StringVar(&text, "text", "", "what `TEXT` tts will say")
	flag.StringVar(&text, "t", "", "what `TEXT` tts will say")
	flag.StringVar(&file, "file", "", "tts read text from `FILE`")
	flag.StringVar(&file, "f", "", "tts read text from `FILE`")
	flag.StringVar(&voice, "voice", "en-US-AriaNeural", "`VOICE` for tts")
	flag.StringVar(&voice, "v", "en-US-AriaNeural", "`VOICE` for tts")
	flag.BoolVar(&listVoices, "list-voices", false, "lists available voices and exits")
	flag.BoolVar(&listVoices, "l", false, "lists available voices and exits")
	flag.StringVar(&rate, "rate", "+0%", "set tts `RATE`")
	flag.StringVar(&volume, "volume", "+0%", "set tts `VOLUME`")
	// flag.StringVar(&wordsInCue, "words-in-cue", "10", "number of `WORDS` in a subtitle cue")
	flag.StringVar(&writeMedia, "write-media", "", "send media output to `FILE` instead of stdout")
	// flag.StringVar(&writeSubtitles, "write-subtitles", "", "send subtitle output to provided `FILE` instead of stderr")
	flag.StringVar(&proxy, "proxy", "", "use a `PROXY` for tts and voice list")
	flag.BoolVar(&fast, "fast", false, "use this flag to run go-playback faster (experimental)")
}

func Parse() {
	flag.Parse()
}

func args() []string {
	return flag.Args()
}

func HasNonFlagArguments() bool {
	return len(args()) > 0
}

func Usage() {
	fmt.Printf("Usage: %[1]s [OPTIONS]\n\n", config.SoftwareName)
	flag.PrintDefaults()
}

func NotACommand() {
	fmt.Printf("%[1]s: '%[2]s' is not a %[1]s command.\n\n", config.SoftwareName, args()[0])
	Usage()
}

func NoExecutableCommand() {
	fmt.Printf("%s: error: one of the arguments -t/-text, -f/-file, -l/-list-voices is required.\n\n", config.SoftwareName)
	Usage()
}

func GetHelp() bool {
	return help
}
func GetText() string {
	return text
}
func GetFile() string {
	return file
}
func GetVoice() string {
	return voice
}
func GetListVoices() bool {
	return listVoices
}
func GetRate() string {
	return rate
}
func GetVolume() string {
	return volume
}
func GetWriteMedia() string {
	return writeMedia
}
func GetProxy() string {
	return proxy
}
func GetFast() bool {
	return fast
}
