package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

type Animation struct {
	anim    *ganim8.Animation
	originX float64
	originY float64
}

type gamePlayer struct {
	x, y         int
	img          *ebiten.Image
	moveDelay    int
	orientation  string
	animInstance Animation
	idle         Animation
	dead         Animation
	walkleft     Animation
	walkright    Animation
	walkup       Animation
	walkdown     Animation
	meleeleft    Animation
	meleeright   Animation
	meleeup      Animation
	meleedown    Animation
}

func NewPlayer() *gamePlayer {
	p := &gamePlayer{
		x:           0,
		y:           0,
		moveDelay:   10, // Adjust this value to change the speed of the player
		orientation: "S",
	}

	return p
}

func LoadPlayerImage(p *gamePlayer) *gamePlayer {
	// p.img, _, _ = ebitenutil.NewImageFromFile("assets/characters/link/link_master.png")
	// melee_img, _, _ := ebitenutil.NewImageFromFile("assets/characters/link/link_sword2.png")
	// idle, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Idle/Idle-Sheet.png")
	// run, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Run/Run-Sheet.png")
	// dead, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Death/Death-Sheet.png")

	p.img, _, _ = ebitenutil.NewImageFromFile("assets/characters/link/link.png")

	g32 := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 0, 0, 1)
	gwalk := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 16, 0, 1)
	gmeleedown := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 4, 0, 1)
	gmeleeup := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 1, 0, 1)

	p.walkleft.anim = ganim8.New(p.img, gwalk.Frames("2", "1-2"), 100*time.Millisecond)
	p.walkleft.originX = 0.35
	p.walkleft.originY = 0.125

	p.walkright.anim = ganim8.New(p.img, gwalk.Frames("4", "1-2"), 10*time.Millisecond)
	p.walkright.originX = 0.1
	p.walkright.originY = 0.1

	p.walkup.anim = ganim8.New(p.img, gwalk.Frames("3", "1-2"), 100*time.Millisecond)
	p.walkup.originX = 0.2
	p.walkup.originY = 0.2

	p.walkdown.anim = ganim8.New(p.img, gwalk.Frames("1", "1-2"), 100*time.Millisecond)
	p.walkdown.originX = 0.4
	p.walkdown.originY = 0.2

	p.meleeright.anim = ganim8.New(p.img, g32.Frames("8", "1"), 100*time.Millisecond)
	p.meleeright.originX = 0.1
	p.meleeright.originY = 0.3

	p.meleeleft.anim = ganim8.New(p.img, g32.Frames("6", "1"), 100*time.Millisecond)
	p.meleeleft.originX = 0.6
	p.meleeleft.originY = 0.3

	p.meleeup.anim = ganim8.New(p.img, gmeleeup.Frames("7", "1"), 100*time.Millisecond)
	p.meleeup.originX = 0.3
	p.meleeup.originY = 0.4

	p.meleedown.anim = ganim8.New(p.img, gmeleedown.Frames("5", "1"), 100*time.Millisecond)
	p.meleedown.originX = 0.4
	p.meleedown.originY = 0.1

	p.animInstance = p.walkdown

	return p
}
