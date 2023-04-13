package airwar2d

import (
	"github.com/LuigiVanacore/AirWars2D/historia_engine"
	"github.com/LuigiVanacore/AirWars2D/historia_engine/input"
	"github.com/LuigiVanacore/AirWars2D/historia_engine/math2d"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	Left = iota
	Right
	Up
)

type Player struct {
	historia_engine.Entity2D
	input.ActionTarget
	velocity  math2d.Vector2D
	direction int
	canMove   bool
}

func NewPlayer() *Player {
	actionMap := input.NewActionMap()
	actionLeft := input.NewActionKey(ebiten.KeyLeft, input.PRESSED)
	actionRight := input.NewActionKey(ebiten.KeyRight, input.PRESSED)
	actionUp := input.NewActionKey(ebiten.KeyUp, input.PRESSED)
	actionMap.Add(1, *actionLeft)
	actionMap.Add(2, *actionRight)
	actionMap.Add(3, *actionUp)
	actionTarget := input.NewActionTarget(actionMap)
	player := &Player{Entity2D: historia_engine.Entity2D{}, ActionTarget: *actionTarget}
	player.SetPosition(150, 150)
	actionTarget.Bind(1, func() { player.canMove = true; player.Move(Left) })
	actionTarget.Bind(2, func() { player.canMove = true; player.Move(Right) })
	actionTarget.Bind(3, func() {
		player.canMove = true
		player.Move(Up)
	})
	return player
}

func (p *Player) Update() {
	p.ProcessEvents()
}

func (p *Player) ProcessEvents() {
	p.canMove = false
	p.ActionTarget.ProcessEvents()
}

func (p *Player) Move(direction int) {
	if p.canMove {
		if direction == Left {
			p.Rotate(-5)
		}
		if direction == Right {
			p.Rotate(5)
		}
		if direction == Up {
			angle := float64(p.GetRotation())/180*math.Pi - math.Pi/2
			p.Transform.Move(math.Cos(angle)*5, math.Sin(angle)*5)
		}
	}
}
