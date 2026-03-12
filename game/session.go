package game

import (
	"encoding/gob"
	"os"

	"github.com/felipelucio/go-pixel-dungeon/core"
)

var session *Session

type Session struct {
	random       *core.Random
	sceneManager *core.SceneManager
	game         *Game
}

type SessionState struct {
	GameState core.GameState
	Seed      string
	SeedState []byte
}

func GetSession() *Session {
	if session == nil {
		panic("Session is not initialized!")
	}
	return session
}

func NewSession(game *Game) *Session {
	s := Session{
		random:       core.NewRandom(),
		game:         game,
		sceneManager: core.NewSceneManager(),
	}
	session = &s
	return &s
}

func LoadSession(game *Game, sessionFile string) (*Session, error) {
	f, err := os.Open(sessionFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var sesState SessionState
	dec := gob.NewDecoder(f)
	err_d := dec.Decode(&sesState)
	if err_d != nil {
		return nil, err
	}

	sess := NewSession(game)
	sess.RNG().LoadState(sesState.Seed, sesState.SeedState)
	sess.Game().SetState(sesState.GameState)

	return sess, nil
}

func (s *Session) SaveSession(sessionFile string) error {
	f, err := os.Create(sessionFile)
	if err != nil {
		return err
	}
	defer f.Close()

	sesState := SessionState{}
	sesState.GameState = s.game.GameState
	seed, st, err := s.RNG().SaveState()
	if err != nil {
		return err
	}
	sesState.Seed = seed
	sesState.SeedState = st
	enc := gob.NewEncoder(f)
	err_e := enc.Encode(&sesState)
	return err_e
}

func (s *Session) RNG() *core.Random {
	return s.random
}

func (s *Session) Game() *Game {
	return s.game
}

func (s *Session) SceneManager() *core.SceneManager {
	return s.sceneManager
}
