package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 600
	screenHeight = 500
)

type Game struct {
	input *Input
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	return &Game{}
}
