package core

type GameState struct {
	HeroName string
	Depth    int
}

func NewGameState(heroName string) *GameState {
	return &GameState{
		HeroName: heroName,
		Depth:    1,
	}
}
