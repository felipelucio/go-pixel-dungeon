package game

type WinModeType int

const (
	Fullscreen WinModeType = iota
	Windowed
	Borderless
)

type config struct {
	AssetsPath string

	WinWidth     int
	WinHeight    int
	WinMode      WinModeType
	VsyncEnabled bool
}

var Config config = config{
	AssetsPath: "assets",

	WinWidth:     768,
	WinHeight:    432,
	WinMode:      Windowed,
	VsyncEnabled: false,
}
