package math

type Vector2Int struct {
	X int
	Y int
}

func NewVector2Int(x int, y int) *Vector2Int {
	return &Vector2Int{X: x, Y: y}
}
