package tts

const (
	trustedClientToken = "6A5AA1D4EAFF4E9FB37E23D68491D6F4"
	WssURL             = "wss://speech.platform.bing.com/consumer/speech/synthesize/" + "readaloud/edge/v1?TrustedClientToken=" + trustedClientToken
	VoiceList          = "https://speech.platform.bing.com/consumer/speech/synthesize/" + "readaloud/voices/list?trustedclienttoken=" + trustedClientToken
)

type Option struct {
	OptID optionID
	Param string
}

type optionID int

const (
	optionIDVoice  optionID = 1
	optionIDRate   optionID = 2
	optionIDVolume optionID = 3
	optionIDProxy  optionID = 4
	optionIDFast   optionID = 5
)

// WithVoice get voice config here: https://learn.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support?tabs=tts
func WithVoice(voice string) Option {
	return Option{
		OptID: optionIDVoice,
		Param: voice,
	}
}

func GetVoiceByOption(opts []Option) string {
	for _, opt := range opts {
		if opt.OptID == optionIDVoice {
			return opt.Param
		}
	}
	return ""
}

func WithRate(rate string) Option {
	return Option{
		OptID: optionIDRate,
		Param: rate,
	}
}

func GetRateByOption(opts []Option) string {
	for _, opt := range opts {
		if opt.OptID == optionIDRate {
			return opt.Param
		}
	}
	return ""
}

func WithVolume(volume string) Option {
	return Option{
		OptID: optionIDVolume,
		Param: volume,
	}
}

func GetVolumeByOption(opts []Option) string {
	for _, opt := range opts {
		if opt.OptID == optionIDVolume {
			return opt.Param
		}
	}
	return ""
}

func WithProxy(proxy string) Option {
	return Option{
		OptID: optionIDProxy,
		Param: proxy,
	}
}

func GetProxyByOption(opts []Option) string {
	for _, opt := range opts {
		if opt.OptID == optionIDProxy {
			return opt.Param
		}
	}
	return ""
}

func WithFast(fast string) Option {
	return Option{
		OptID: optionIDFast,
		Param: fast,
	}
}

func GetFastByOption(opts []Option) string {
	for _, opt := range opts {
		if opt.OptID == optionIDFast {
			return opt.Param
		}
	}
	return ""
}
