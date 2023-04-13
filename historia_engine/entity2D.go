package historia_engine

type Entity2D struct {
	Transform
}

func (e *Entity2D) Update() {
}

func (e *Entity2D) GetTransform() Transform {
	return e.Transform
}

func (e *Entity2D) SetTransform(transform Transform) {
	e.Transform = transform
}
