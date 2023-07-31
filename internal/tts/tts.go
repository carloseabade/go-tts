package tts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/carloseabade/go-tts/internal/config"
	"github.com/carloseabade/go-tts/pkg/tts"
	file_helper "github.com/pp-group/file-helper"
	storage "github.com/pp-group/file-helper/storage"
)

type ISpeech interface {
	GenTTS() error
	URL(filename string) (string, error)
}

var _ ISpeech = new(LocalSpeech)

type LocalSpeech struct {
	*Speech
}

func NewLocalSpeech(c *tts.Communicate, folder, filename string) (*LocalSpeech, error) {

	fileStorage, err := file_helper.FileStorageFactory(folder)()
	if err != nil {
		return nil, err
	}

	s, err := NewSpeech(c, fileStorage, folder, filename)
	if err != nil {
		return nil, err
	}

	return &LocalSpeech{
		Speech: s,
	}, nil

}

func (speech *LocalSpeech) GenTTS() error {
	return gentts(speech.Speech, func() (storage.IWriteBroker, error) {
		return speech.Writer(speech.FileName, nil)
	})
}

func (speech *LocalSpeech) URL(filename string) (string, error) {
	return urlTTS(func() (storage.IReadBroker, error) {
		return speech.Reader(filename, nil)
	})
}

var _ ISpeech = new(OssSpeech)

type OssSpeech struct {
	*Speech
	bucket string
}

func NewOssSpeech(c *tts.Communicate, endpoint, ak, sk, folder, bucket string) (*OssSpeech, error) {

	ossStorage, err := file_helper.OssStorageFactory(endpoint, ak, sk, folder)()
	if err != nil {
		return nil, err
	}

	s, err := NewSpeech(c, ossStorage, folder, "")
	if err != nil {
		return nil, err
	}

	return &OssSpeech{
		Speech: s,
		bucket: bucket,
	}, nil
}

func (speech *OssSpeech) GenTTS() error {

	return gentts(speech.Speech, func() (storage.IWriteBroker, error) {
		return speech.Writer(speech.FileName, func() interface{} {
			return speech.bucket
		})
	})
}

func (speech *OssSpeech) URL(filename string) (string, error) {
	return urlTTS(func() (storage.IReadBroker, error) {
		return speech.Reader(filename, func() interface{} {
			return speech.bucket
		})
	})
}
func gentts(speech *Speech, brokerFunc func() (storage.IWriteBroker, error)) error {

	if speech.FileName == "" {
		speech.FileName = generateHashName(speech.Text, speech.VoiceLangRegion) + ".mp3"
	}

	broker, err := brokerFunc()
	if err != nil {
		return err
	}

	err = speech.gen(broker)
	if err != nil {
		return err
	}
	return nil
}

func urlTTS(brokerFunc func() (storage.IReadBroker, error)) (string, error) {

	broker, err := brokerFunc()
	if err != nil {
		return "", err
	}
	return broker.URL()
}

type Speech struct {
	*tts.Communicate
	storage.IStorage
	Folder   string
	FileName string
}

func NewSpeech(c *tts.Communicate, storage storage.IStorage, folder, filename string) (*Speech, error) {
	s := &Speech{
		Communicate: c,
		IStorage:    storage,
		Folder:      folder,
		FileName:    filename,
	}
	return s, nil

}

func (s *Speech) gen(broker storage.IWriteBroker) error {
	op, err := s.Stream()
	if err != nil {
		return err
	}
	defer s.CloseOutput()
	solveCount := 0
	audioData := make([][][]byte, s.AudioDataIndex)
	for i := range op {
		if _, ok := i["end"]; ok {
			solveCount++
			if solveCount == s.AudioDataIndex {
				break
			}
		}
		t, ok := i["type"]
		if ok && t == "audio" {
			data := i["data"].(tts.AudioData)
			audioData[data.Index] = append(audioData[data.Index], data.Data)
		}
		e, ok := i["error"]
		if ok {
			fmt.Printf("has error err: %v\n", e)
		}
	}
	// write data, sort by index
	for _, v := range audioData {
		for _, data := range v {
			broker.Write(data)
		}
	}
	broker.Close()
	return nil
}

func generateHashName(name, voice string) string {
	hash := sha256.Sum256([]byte(name))
	return fmt.Sprintf("%s_%s", voice, hex.EncodeToString(hash[:]))
}

type OssSpeechFactory struct {
	endpoint string
	ak       string
	sk       string
	bucket   string
	folder   string
}

func NewOssSpeechFactory(endpoint, ak, sk, bucket, folder string) *OssSpeechFactory {
	return &OssSpeechFactory{
		endpoint: endpoint,
		ak:       ak,
		sk:       sk,
		bucket:   bucket,
		folder:   folder,
	}
}

func (factory *OssSpeechFactory) OssSpeech(c *tts.Communicate, folder string) (*OssSpeech, error) {
	if folder != "" {
		return NewOssSpeech(c, factory.endpoint, factory.ak, factory.sk, folder, factory.bucket)
	}
	return NewOssSpeech(c, factory.endpoint, factory.ak, factory.sk, factory.folder, factory.bucket)
}

type TTSMode int

const (
	TEXT TTSMode = 0
	FILE TTSMode = 1
)

func TTS(f TTSMode, input, writeMedia, voice, rate, volume, proxy string) {
	if f == TEXT {
		TTSText(input, writeMedia, tts.WithVoice(voice), tts.WithRate(rate), tts.WithVolume(volume), tts.WithProxy(proxy))
	} else if f == FILE {
		TTSFile(input, writeMedia, tts.WithVoice(voice), tts.WithRate(rate), tts.WithVolume(volume), tts.WithProxy(proxy))
	} else {
		handleError(fmt.Errorf("TTS function internal error"))
	}
}

func TTSText(text, writeMedia string, opts ...tts.Option) {
	c, err := tts.NewCommunicate(text, opts...)
	handleError(err)

	speech, err := NewLocalSpeech(c, "", writeMedia)
	handleError(err)

	err = speech.GenTTS()
}

func TTSFile(file, writeMedia string, opts ...tts.Option) {
	dat, err := os.ReadFile(file)
	handleError(err)

	TTSText(string(dat), writeMedia, opts...)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", config.SoftwareName, err)
		os.Exit(1)
	}
}
