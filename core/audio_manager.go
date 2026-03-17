package core

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const (
	SAMPLE_RATE = 48000
)

type AudioManager struct {
	context *audio.Context

	soundsOgg map[string]*vorbis.Stream
	musicOgg  map[string]*vorbis.Stream
	soundsMp3 map[string]*mp3.Stream
	musicMp3  map[string]*mp3.Stream
	players   map[string]chan *audio.Player
}

func NewAudioManager() *AudioManager {
	ctx := audio.NewContext(SAMPLE_RATE)
	return &AudioManager{
		context:   ctx,
		soundsOgg: make(map[string]*vorbis.Stream),
		musicOgg:  make(map[string]*vorbis.Stream),
		soundsMp3: make(map[string]*mp3.Stream),
		musicMp3:  make(map[string]*mp3.Stream),
		players:   make(map[string]chan *audio.Player),
	}
}

func (am *AudioManager) LoadMusic(mList map[string]string) error {
	for name, path := range mList {
		p := fmt.Sprintf("%s/%s", "assets", path)
		ext := filepath.Ext(p)
		switch ext {
		case ".ogg":
			m, err := loadOGG(p, SAMPLE_RATE)
			if err != nil {
				fmt.Println(err)
			}
			am.musicOgg[name] = m
			return nil
		case ".mp3":
			m, err := loadMP3(p, SAMPLE_RATE)
			if err != nil {
				fmt.Println(err)
			}
			am.musicMp3[name] = m
			return nil
		default:
			return fmt.Errorf("Wrong file type: %s", p)
		}
	}
	return nil
}

func (am *AudioManager) LoadSFX(mList map[string]string) error {
	for name, path := range mList {
		p := fmt.Sprintf("%s/%s", "assets", path)
		ext := filepath.Ext(p)
		switch ext {
		case ".ogg":
			m, err := loadOGG(p, SAMPLE_RATE)
			if err != nil {
				fmt.Println(err)
			}
			am.soundsOgg[name] = m
			return nil
		case ".mp3":
			m, err := loadMP3(p, SAMPLE_RATE)
			if err != nil {
				fmt.Println(err)
			}
			am.soundsMp3[name] = m
			return nil
		default:
			return fmt.Errorf("Wrong file type: %s", p)
		}
	}
	return nil
}

func (am *AudioManager) PlaySFX(name string) {

}

func (am *AudioManager) PlayMusic(name string) {

}

func loadOGG(filePath string, sampleRate int) (*vorbis.Stream, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	stream, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func loadMP3(filePath string, sampleRate int) (*mp3.Stream, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	stream, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	return stream, nil
}
