package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1000
	screenHeight = 800
)

type Game struct {
	input *Input
	man   *Man
}

func (g *Game) Update() error {
	g.input.Update()
	g.man.Update(g.input)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.man.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	i := NewInput()
	m, err := NewMan()
	if err != nil {
		panic(err)
	}
	return &Game{
		input: i,
		man:   m,
	}
}
