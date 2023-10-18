package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1000
	screenHeight = 800
)

type Game struct {
	input    *Input
	man      *Man
	obstacle *Obstacle
}

func (g *Game) Update() error {
	g.input.Update()
	g.man.Update(g.input, g.obstacle)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.man.Draw(screen)
	g.obstacle.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	i := NewInput()
	m, err1 := NewMan()
	if err1 != nil {
		panic(err1)
	}
	o, err2 := NewObstacle()
	if err2 != nil {
		panic(err2)
	}
	return &Game{
		input:    i,
		man:      m,
		obstacle: o,
	}
}
