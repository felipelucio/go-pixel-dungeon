package core

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var _defaultAssetManager *AssetManager

type AssetManager struct {
	// effects map[string]*core.Effect
	fonts  map[string]*text.GoTextFaceSource
	images map[string]*ebiten.Image
	audios map[string]io.ReadSeeker
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		fonts:  make(map[string]*text.GoTextFaceSource),
		images: make(map[string]*ebiten.Image),
		audios: make(map[string]io.ReadSeeker),
	}
}

func DefaultAssetManager() *AssetManager {
	if _defaultAssetManager == nil {
		_defaultAssetManager = NewAssetManager()
	}
	return _defaultAssetManager
}

func SetDefaultAssetManager(am *AssetManager) {
	_defaultAssetManager = am
}

func (am *AssetManager) LoadFont(name string, filePath string) (*text.GoTextFaceSource, error) {
	font, ok := am.fonts[name]
	if ok {
		return font, nil
	}

	fpath := path.Join("./assets", filePath)
	ext := filepath.Ext(fpath)
	buf, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	switch ext {
	case ".ttf":
		font, err := text.NewGoTextFaceSource(bytes.NewReader(buf))
		if err != nil {
			return nil, err
		}
		am.fonts[name] = font
		return font, nil
	default:
		return nil, fmt.Errorf("Wrong file type '%s' in '%s'", ext, filePath)
	}
}

func (am *AssetManager) LoadFontList(fontList map[string]string) {
	for name, fpath := range fontList {
		_, err := am.LoadFont(name, fpath)
		if err != nil {
			slog.Warn("[AssetManager] Font not loaded", "file", fpath, "error", err)
		}
	}
}

func (am *AssetManager) LoadImage(name string, filePath string) (*ebiten.Image, error) {
	img, ok := am.images[name]
	if ok {
		return img, nil
	}

	fpath := path.Join("./assets", filePath)
	buf, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	rawimg, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	img = ebiten.NewImageFromImage(rawimg)
	am.images[name] = img
	return img, nil
}

func (am *AssetManager) LoadImageList(fileList map[string]string) {
	for name, fpath := range fileList {
		_, err := am.LoadImage(name, fpath)
		if err != nil {
			slog.Warn("[AssetManager] Image not loaded", "file", fpath, "error", err)
		}
	}
}

func (am *AssetManager) LoadAudio(name string, filePath string) (io.ReadSeeker, error) {
	aud, ok := am.audios[name]
	if ok {
		return aud, nil
	}
	var err error
	fpath := path.Join("./assets", filePath)
	ext := filepath.Ext(fpath)
	buf, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	var stream io.ReadSeeker
	switch ext {
	case ".mp3":
		stream, err = mp3.DecodeF32(buf)
	case ".ogg":
		stream, err = vorbis.DecodeF32(buf)
	case ".wav":
		stream, err = wav.DecodeF32(buf)
	default:
		return nil, fmt.Errorf("Audio format '%s' is invalid in '%s'", ext, filePath)
	}
	if err != nil {
		return nil, err
	}
	am.audios[name] = stream

	return stream, nil
}

func (am *AssetManager) LoadAudioList(fileList map[string]string) {
	for name, fpath := range fileList {
		_, err := am.LoadAudio(name, fpath)
		if err != nil {
			slog.Warn("[AssetManager] Audio not loaded", "file", fpath, "error", err)
		}
	}
}

func (am *AssetManager) GetAudio(name string) (io.ReadSeeker, bool) {
	au, ok := am.audios[name]
	return au, ok
}

func (am *AssetManager) GetImage(name string) (*ebiten.Image, bool) {
	au, ok := am.images[name]
	return au, ok
}

func (am *AssetManager) GetFont(name string) (*text.GoTextFaceSource, bool) {
	au, ok := am.fonts[name]
	return au, ok
}
