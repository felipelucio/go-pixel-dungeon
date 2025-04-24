package core

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init(state *GameState) error
	Pause()
	Resume()
	Update() error
	Draw(screen *ebiten.Image) error
	Destroy()
}

var (
	sceneStack = make([]Scene, 5)
	gameState  *GameState
)

func SceneManagerRegisterGameState(gs *GameState) {
	gameState = gs
}

// Scene returns the current scene.
func CurrScene() Scene {
	return sceneStack[len(sceneStack)-1]
}

func SwitchToScene(scene Scene) {
	currentScene := CurrScene()
	for currentScene != nil {
		currentScene.Destroy()
	}
	sceneStack = sceneStack[:0]
	sceneStack = append(sceneStack, scene)
	scene.Init(gameState)
	scene.Resume()
}

func PushScene(scene Scene) {
	currentScene := CurrScene()
	for currentScene != nil {
		currentScene.Pause()
	}
	scene.Init(gameState)
	sceneStack = append(sceneStack, scene)
	scene.Resume()
}

// Update is the function called every frame by ebiten.
func UpdateScene() error {
	currentScene := CurrScene()
	if currentScene == nil {
		return nil
	}

	return currentScene.Update()
}

func DrawScene(screen *ebiten.Image) error {
	currentScene := CurrScene()
	if currentScene == nil {
		return nil
	}
	return currentScene.Draw(screen)
}
