# go-tts

`go-tts` is a go module that allows you to use Microsoft Edge's online text-to-speech service. It is inspired on [edge-tts](https://github.com/rany2/edge-tts) and forked from [edge-tts-go](https://github.com/rany2/edge-tts).

## Usage

### Basic usage

To use the `go-tts`, run the following:

```bash
$ go-tts -text "Hello, world!" -write-media hello.mp3
```

If you wish to play it back immediately, use the `go-playback` command:

```bash
$ go-playback -text "Hello, world!"
```

Note the above requires the installation of the `mpv` command line player.

**NOTE**: All go-tts commands work in go-playback as well.

### Changing the voice

If you want to change the voice language, run the following: 

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

It is possible to make minor changes to the generated speech.

```bash
$ go-tts -rate -50% -text "Hello, world!" -write-media hello_with_rate_halved.mp3
$ go-tts -volume -50% -text "Hello, world!" -write-media hello_with_volume_halved.mp3
```

## LICENSE: MIT
github.com/carloseabade/go-tts is licensed under the **MIT License**
