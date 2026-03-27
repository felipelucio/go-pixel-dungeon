package game

import (
	"encoding/gob"
	"os"

	"github.com/felipelucio/go-pixel-dungeon/core"
)

var _currSession *Session

type Session struct {
	Random    *core.Random
	GameState GameState
	Seed      string
}

// For serialization only
type sessionState struct {
	GameState GameState
	Seed      string
	SeedState []byte
}

func GetSession() *Session {
	if _currSession == nil {
		panic("Session is not initialized!")
	}
	return _currSession
}

func SetSession(sess *Session) {
	_currSession = sess
}

func NewSession(seed *string) *Session {
	rng := core.NewRandom(seed)
	// rng.Load
	s := &Session{
		Random:    rng,
		GameState: *NewGameState(),
		Seed:      rng.GetSeed(),
	}
	return s
}

func LoadSession(sessionFile string) (*Session, error) {
	f, err := os.Open(sessionFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sesState sessionState
	dec := gob.NewDecoder(f)
	err_d := dec.Decode(&sesState)
	if err_d != nil {
		return nil, err
	}

	sess := NewSession(nil)
	sess.RNG().LoadState(sesState.Seed, sesState.SeedState)

	return sess, nil
}

func (s *Session) SaveSession(sessionFile string) error {
	f, err := os.Create(sessionFile)
	if err != nil {
		return err
	}
	defer f.Close()

	sesState := sessionState{}
	seed, st, err := s.RNG().SaveState()
	if err != nil {
		return err
	}
	sesState.Seed = seed
	sesState.SeedState = st
	sesState.GameState = s.GameState
	enc := gob.NewEncoder(f)
	err_e := enc.Encode(&sesState)
	return err_e
}

func (s *Session) RNG() *core.Random {
	return s.Random
}
