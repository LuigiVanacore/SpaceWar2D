package airwar2d

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

type Game struct {
	ready bool
	world *World
}

func (g *Game) Init() {
	g.world = NewWorld()
}

func (g *Game) HandleInput() error {

	return nil
}

func (g *Game) Update() error {
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.camera.Surface.Clear()
	//g.camera.Draw(g.world)
	//g.camera.Blit(screen)
	op := &ebiten.DrawImageOptions{}
	g.world.Draw(screen, op)
	//playerPosition := g.world.GetPlayer().GetPosition()
	//transform := g.world.GetPlayer().GetTransform()
	//x, y := transform.GetPivot()
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("player %3.3f %3.3f \n pivot %3.3f %3.3f \n", playerPosition.X, playerPosition.Y, x, y))

}

func (g *Game) Layout(int, int) (screenwidth int, screenheight int) {
	return WindowWidth, WindowHeight
}
