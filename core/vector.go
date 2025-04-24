package core

type Vector2[T any] struct {
	X T
	Y T
}

func NewVector2[T any](x T, y T) Vector2[T] {
	return Vector2[T]{
		x,
		y,
	}
}

type Vector3[T any] struct {
	X T
	Y T
	Z T
}

func NewVector3[T any](x T, y T, z T) Vector3[T] {
	return Vector3[T]{
		x,
		y,
		z,
	}
}
