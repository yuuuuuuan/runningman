package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

var (
	runningImg *ebiten.Image
)

const (
	movingSpeed = 1
)

type Frame struct {
	x      int
	y      int
	width  int
	height int
}

type Man struct {
	frame Frame
	img   *ebiten.Image
}

func (m *Man) IfCollision(o *Obstacle) (dir, bool) {
	return CheckCollision(m.frame, o.frame)
}

func NewMan() (*Man, error) {
	img, _, errImg := ebitenutil.NewImageFromFile("D:\\Program Files\\JetBrains\\GoLand 2023.1.3\\workspaces\\runningman\\game\\img.png")
	if errImg != nil {
		panic(errImg)
	}
	runningImg = ebiten.NewImageFromImage(img)
	w := runningImg.Bounds().Dx()
	h := runningImg.Bounds().Dy()
	f := &Frame{
		x:      0,
		y:      0,
		width:  w,
		height: h,
	}
	m := &Man{
		frame: *f,
		img:   runningImg,
	}
	return m, errImg
}

func (m *Man) Update(i *Input, o *Obstacle) {
	dir, _ := m.IfCollision(o)

	switch i.state {
	case keyup:
		if dir == "down" {
			return
		}
		m.frame.y -= movingSpeed
	case keydown:
		if dir == "up" {
			return
		}
		m.frame.y += movingSpeed
	case keyleft:
		if dir == "right" {
			return
		}
		m.frame.x -= movingSpeed
	case keyright:
		if dir == "left" {
			return
		}
		m.frame.x += movingSpeed
	}
}

func (m *Man) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(m.frame.x), float64(m.frame.y))
	screen.DrawImage(m.img, op)
}
