package airwar2d

import (
	"fmt"
	"github.com/LuigiVanacore/AirWars2D/Assets"
	historia "github.com/LuigiVanacore/AirWars2D/historia_engine"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

type Game struct {
	ready  bool
	world  *historia.SceneNode
	player *Player
	desert *historia.SpriteNode
	camera *historia.Camera
}

func (g *Game) Init() {
	g.camera = historia.NewCamera(WindowWidth, WindowHeight, WindowWidth/2, WindowHeight/2, 0, 1)
	g.world = historia.NewSceneNode()
	g.player = NewPlayer()
	g.desert = historia.NewSpriteNode(Assets.ResourceManager().GetTexture(Assets.Desert))
	g.camera.Follow(g.player)
	//g.world.AttachChild(g.desert)
	g.world.AttachChild(g.player)

}

func (g *Game) HandleInput() error {

	return nil
}

func (g *Game) Update() error {
	g.world.Update()
	g.camera.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.camera.Surface.Clear()
	//g.camera.Draw(g.world)
	//g.camera.Blit(screen)
	op := &ebiten.DrawImageOptions{}
	g.world.Draw(screen, op)
	x, _ := g.player.GetPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("player %3.3f \n  camera %3.3f", x, g.camera.X))

}

func (g *Game) Layout(outsidewidth, outsidheight int) (screenwidth, screenheight int) {
	return WindowWidth, WindowHeight
}
