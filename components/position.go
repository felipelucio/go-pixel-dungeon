package components

type Position struct {
	X int32
	Y int32
}

func (pos Position) GetName() string {
	return "Position"
}
