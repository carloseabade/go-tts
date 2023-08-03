# go-tts

`go-tts` is a Go module that enables you to utilize Microsoft Edge's online text-to-speech service. It is inspired by [edge-tts](https://github.com/rany2/edge-tts) and forked from [edge-tts-go](https://github.com/rany2/edge-tts).

## Installation

You can download and install the packages running:

```bash
$ go install github.com/carloseabade/go-tts/cmd/go-tts@latest
$ go install github.com/carloseabade/go-tts/cmd/go-playback@latest
```

## Usage

### Basic usage

To use `go-tts`, execute the following:

```bash
$ go-tts -text "Hello, world!" -write-media hello.mp3
```

**NOTE**: You can always use `-help` to learn all the available options :)

### Changing the voice

If you wish to change the voice language, follow these steps:

You can check the available voices with the `-list-voices` option:

```bash
$ go-tts -list-voices
...
Name: Microsoft Server Speech Text to Speech Voice (ps-AF, GulNawazNeural)
ShortName: ps-AF-GulNawazNeural
Gender: Male
Locale: ps-AF

Name: Microsoft Server Speech Text to Speech Voice (ps-AF, LatifaNeural)
ShortName: ps-AF-LatifaNeural
Gender: Female
Locale: ps-AF

Name: Microsoft Server Speech Text to Speech Voice (pt-BR, AntonioNeural)
ShortName: pt-BR-AntonioNeural
Gender: Male
Locale: pt-BR

Name: Microsoft Server Speech Text to Speech Voice (pt-BR, FranciscaNeural)
ShortName: pt-BR-FranciscaNeural
Gender: Female
Locale: pt-BR
...
$ go-tts -voice pt-BR-AntonioNeural -text "Olá, mundo!" -write-media hello_in_portuguese.mp3
```

### Changing rate and volume

You can make minor adjustments to the generated speech.

```bash
$ go-tts -rate -50% -text "Hello, world!" -write-media hello_with_rate_halved.mp3
$ go-tts -volume -50% -text "Hello, world!" -write-media hello_with_volume_halved.mp3
```

## LICENSE: MIT
github.com/carloseabade/go-tts is licensed under the **MIT License**