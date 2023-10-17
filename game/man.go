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
	movingSpeed = 3
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

func NewMan() (*Man, error) {
	img, _, errImg := ebitenutil.NewImageFromFile("D:\\Program Files\\JetBrains\\GoLand 2023.1.3\\workspaces\\runningman\\game\\img.png")
	if errImg != nil {
		panic(errImg)
	}
	runningImg = ebiten.NewImageFromImage(img)
	w := runningImg.Bounds().Dx()
	h := runningImg.Bounds().Dy()
	f := &Frame{
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		width:  w,
		height: h,
	}
	m := &Man{
		frame: *f,
		img:   runningImg,
	}
	return m, errImg
}

func (m *Man) Update(i *Input) {
	switch i.state {
	case keyup:
		m.frame.y -= movingSpeed
	case keydown:
		m.frame.y += movingSpeed
	case keyleft:
		m.frame.x -= movingSpeed
	case keyright:
		m.frame.x += movingSpeed
	}
}

func (m *Man) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(m.frame.x), float64(m.frame.y))
	screen.DrawImage(m.img, op)
}
