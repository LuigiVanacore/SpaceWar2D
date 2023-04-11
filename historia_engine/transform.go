package historia_engine

import "github.com/LuigiVanacore/AirWars2D/historia_engine/math2d"

type Transform struct {
	position math2d.Vector2D
	pivot    math2d.Vector2D
	rotation int
}

func (t *Transform) GetPosition() (float64, float64) {
	return t.position.X, t.position.Y
}

func (t *Transform) SetPosition(x, y float64) {
	t.position.X = x
	t.position.Y = y
}

func (t *Transform) GetRotation() int {
	return t.rotation
}

func (t *Transform) SetRotation(rotation int) {
	t.rotation = rotation
}

func (t *Transform) SetPivot(x, y float64) {
	t.pivot.X = x
	t.pivot.Y = y
}

func (t *Transform) GetPivot() (float64, float64) {
	return t.pivot.X, t.pivot.Y
}

func (t *Transform) Rotate(rotation int) {
	t.rotation += rotation
}

func (t *Transform) Move(x, y float64) {
	t.position.X += x
	t.position.Y += y
}
