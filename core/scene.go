package core

import "github.com/hajimehoshi/ebiten/v2"

var _defaultSceneManager *SceneManager

type Scene interface {
	Init() error
	Pause()
	Resume()
	Update() error
	Draw(screen *ebiten.Image) error
	Destroy()
}

type SceneManager struct {
	sceneStack []Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		sceneStack: make([]Scene, 5),
	}
}

func DefaultSceneManager() *SceneManager {
	if _defaultSceneManager == nil {
		_defaultSceneManager = NewSceneManager()
	}
	return _defaultSceneManager
}

func SetDefaultSceneManager(sm *SceneManager) {
	_defaultSceneManager = sm
}

// Scene returns the current scene.
func (sm *SceneManager) CurrScene() Scene {
	return sm.sceneStack[len(sm.sceneStack)-1]
}

func (sm *SceneManager) SwitchToScene(scene Scene) {
	currentScene := sm.CurrScene()
	for currentScene != nil {
		currentScene.Destroy()
	}
	sm.sceneStack = sm.sceneStack[:0]
	sm.sceneStack = append(sm.sceneStack, scene)
	scene.Init()
	scene.Resume()
}

func (sm *SceneManager) PushScene(scene Scene) {
	currentScene := sm.CurrScene()
	for currentScene != nil {
		currentScene.Pause()
	}
	scene.Init()
	sm.sceneStack = append(sm.sceneStack, scene)
	scene.Resume()
}

// Update is the function called every frame by ebiten.
func (sm *SceneManager) UpdateScene() error {
	currentScene := sm.CurrScene()
	if currentScene == nil {
		return nil
	}

	return currentScene.Update()
}

func (sm *SceneManager) DrawScene(screen *ebiten.Image) error {
	currentScene := sm.CurrScene()
	if currentScene == nil {
		return nil
	}
	return currentScene.Draw(screen)
}
