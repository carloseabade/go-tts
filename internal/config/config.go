package config

type MessageType int

const (
	SoftwareName       = "go-tts"
	TrustedClientToken = "6A5AA1D4EAFF4E9FB37E23D68491D6F4"
	WssURL             = "wss://speech.platform.bing.com/consumer/speech/synthesize/" + "readaloud/edge/v1?TrustedClientToken=" + TrustedClientToken
	VoiceList          = "https://speech.platform.bing.com/consumer/speech/synthesize/" + "readaloud/voices/list?trustedclienttoken=" + TrustedClientToken
	DefaultVoice       = "en-US-AriaNeural"
)
