package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

type gamePlayer struct {
	x, y         int
	img          *ebiten.Image
	moveDelay    int
	animInstance *ganim8.Animation
	walkx        *ganim8.Animation
	walky        *ganim8.Animation
	melee        *ganim8.Animation
}

func NewPlayer() *gamePlayer {
	p := &gamePlayer{
		x:         0,
		y:         0,
		moveDelay: 10, // Adjust this value to change the speed of the player

	}

	return p
}

func LoadPlayerImage(p *gamePlayer) *gamePlayer {
	// p.img, _, _ = ebitenutil.NewImageFromFile("assets/characters/link/link_master.png")
	// melee_img, _, _ := ebitenutil.NewImageFromFile("assets/characters/link/link_sword2.png")
	p.img, _, _ = ebitenutil.NewImageFromFile("assets/characters/link/link.png")

	// g32 := ganim8.NewGrid(32, 32, 2580, 617)
	// g16 := ganim8.NewGrid(32, 32, 128, 32)
	// p.walkx = ganim8.New(p.img, g32.Frames("1-5", 6), 100*time.Millisecond)
	// p.walky = ganim8.New(p.img, g32.Frames("1-5", 5), 100*time.Millisecond)
	// p.melee = ganim8.New(melee_img, g16.Frames("4-1", 1), 100*time.Millisecond)
	// p.animInstance = ganim8.New(p.img, g32.Frames("1-5", 6), 100*time.Millisecond)

	g32 := ganim8.NewGrid(16, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 1, 1, 1)
	p.walkx = ganim8.New(p.img, g32.Frames("8-10", "2-3"), 100*time.Millisecond)
	p.walky = ganim8.New(p.img, g32.Frames("4-6", "2-3"), 100*time.Millisecond)
	p.melee = ganim8.New(p.img, g32.Frames("12-14", "2-3"), 100*time.Millisecond)
	p.animInstance = p.walkx

	return p
}
