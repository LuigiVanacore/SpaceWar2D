package historia_engine

type Transformable interface {
	GetTransform() Transform
	SetTransform(transform Transform)
}
