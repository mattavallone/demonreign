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
	orientation  string
	animInstance *ganim8.Animation
	idle         *ganim8.Animation
	dead         *ganim8.Animation
	walkleft     *ganim8.Animation
	walkright    *ganim8.Animation
	walkup       *ganim8.Animation
	walkdown     *ganim8.Animation
	meleeleft    *ganim8.Animation
	meleeright   *ganim8.Animation
	meleeup      *ganim8.Animation
	meleedown    *ganim8.Animation
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
	// idle, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Idle/Idle-Sheet.png")
	// run, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Run/Run-Sheet.png")
	// dead, _, _ := ebitenutil.NewImageFromFile("assets/PixelCrawler/Heroes/Rogue/Death/Death-Sheet.png")

	p.img, _, _ = ebitenutil.NewImageFromFile("assets/characters/link/link.png")

	g32 := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 0, 0, 1)
	gwalk := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 16, 0, 1)
	gmeleedown := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 4, 0, 1)
	gmeleeup := ganim8.NewGrid(32, 32, p.img.Bounds().Dx(), p.img.Bounds().Dy(), 1, 0, 1)

	p.walkleft = ganim8.New(p.img, gwalk.Frames("2", "1-2"), 100*time.Millisecond)
	p.walkright = ganim8.New(p.img, gwalk.Frames("4", "1-2"), 10*time.Millisecond)
	p.walkup = ganim8.New(p.img, gwalk.Frames("3", "1-2"), 100*time.Millisecond)
	p.walkdown = ganim8.New(p.img, gwalk.Frames("1", "1-2"), 100*time.Millisecond)
	p.meleeright = ganim8.New(p.img, g32.Frames("8", "1"), 100*time.Millisecond)
	p.meleeleft = ganim8.New(p.img, g32.Frames("6", "1"), 100*time.Millisecond)
	p.meleeup = ganim8.New(p.img, gmeleeup.Frames("7", "1"), 100*time.Millisecond)
	p.meleedown = ganim8.New(p.img, gmeleedown.Frames("5", "1"), 100*time.Millisecond)
	p.animInstance = p.walkright

	return p
}
