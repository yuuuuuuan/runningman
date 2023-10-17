package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	framemanWidth  = 60
	framemanHeight = 100
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

func NewMan() *Man {
	img, _, err := ebitenutil.NewImageFromReader("/img.png")
	w, h := img.Size(img)
	f := &Frame{
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		width:  w,
		height: h,
	}
	m := &Man{
		frame: *f,
		img:   img,
	}
	return m
}
