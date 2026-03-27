package game

type GameState struct {
	HeroName string
	Depth    int
}

func NewGameState() *GameState {
	return &GameState{
		HeroName: "heroName",
		Depth:    1,
	}
}
