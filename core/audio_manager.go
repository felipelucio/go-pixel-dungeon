package core

import (
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	SAMPLE_RATE = 48000
)

var _defaultAudioManager *AudioManager

type AudioManager struct {
	context      *audio.Context
	sfxStreams   map[string]io.ReadSeeker
	sfxPlayers   []*audio.Player
	musicStreams map[string]*audio.InfiniteLoop
	musicPlayer  *audio.Player
}

func NewAudioManager() *AudioManager {
	ctx := audio.NewContext(SAMPLE_RATE)
	return &AudioManager{
		context:      ctx,
		sfxStreams:   make(map[string]io.ReadSeeker),
		sfxPlayers:   make([]*audio.Player, 10),
		musicStreams: make(map[string]*audio.InfiniteLoop),
	}
}

func DefaultAudioManager() *AudioManager {
	if _defaultAudioManager == nil {
		_defaultAudioManager = NewAudioManager()
	}
	return _defaultAudioManager
}

func SetDefaultAudioManager(am *AudioManager) {
	_defaultAudioManager = am
}

func (am *AudioManager) RegisterSFX(name string, stream io.ReadSeeker) {
	am.sfxStreams[name] = stream
}

func (am *AudioManager) RegisterMusic(name string, stream io.ReadSeeker) error {
	sz, err := stream.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	_, _ = stream.Seek(0, io.SeekStart)
	s := audio.NewInfiniteLoopF32(stream, sz)
	am.musicStreams[name] = s
	return nil
}

func (am *AudioManager) PlaySFX(name string) (*audio.Player, error) {
	stream, ok := am.sfxStreams[name]
	if !ok {
		return nil, fmt.Errorf("SFX not registered: %s", name)
	}
	pl, err := am.context.NewPlayerF32(stream)
	if err != nil {
		return nil, err
	}
	am.sfxPlayers = append(am.sfxPlayers, pl)
	pl.Play()
	return pl, nil
}

func (am *AudioManager) PlayMusic(name string) (*audio.Player, error) {
	stream, ok := am.sfxStreams[name]
	if !ok {
		return nil, fmt.Errorf("Music not registered: %s", name)
	}
	pl, err := am.context.NewPlayerF32(stream)
	if err != nil {
		return nil, err
	}
	am.musicPlayer = pl
	pl.Play()
	return pl, nil
}
