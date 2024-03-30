package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

type gamePrincess struct {
	x, y        int
	moveDelay   int
	orientation string
	health      int

	animInstance Animation
}

func LoadPrincess() *gamePrincess {
	p := &gamePrincess{
		x:           5,
		y:           4,
		health:      10,
		moveDelay:   10,
		orientation: "S",
	}

	return p
}

func LoadPrincessImage(p *gamePrincess) *gamePrincess {
	princessImg, _, _ := ebitenutil.NewImageFromFile("assets/characters/princess.png")

	g32 := ganim8.NewGrid(300, 300, princessImg.Bounds().Dx(), princessImg.Bounds().Dy(), 0, 0, 1)

	p.animInstance.anim = ganim8.New(princessImg, g32.Frames(1, 1), 10*time.Millisecond)
	p.animInstance.originX = 0
	p.animInstance.originY = 0

	return p
}
