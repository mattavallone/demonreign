package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	tileSize     = 32
	moveDelay    = 10 // Adjust this value to change the speed of the player
)

var (
	playerX      = 0
	playerY      = 0
	mapSize      = 10
	gameMap      = make([][]int, mapSize)
	moveCooldown = 0
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range gameMap {
		gameMap[i] = make([]int, mapSize)
		for j := range gameMap[i] {
			gameMap[i][j] = rand.Intn(2)
		}
	}
	playerX = rand.Intn(mapSize)
	playerY = rand.Intn(mapSize)
}

func update(screen *ebiten.Image) error {
	if moveCooldown > 0 {
		moveCooldown--
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if playerY > 0 {
			playerY--
		}
		moveCooldown = moveDelay
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if playerY < mapSize-1 {
			playerY++
		}
		moveCooldown = moveDelay
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if playerX > 0 {
			playerX--
		}
		moveCooldown = moveDelay
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if playerX < mapSize-1 {
			playerX++
		}
		moveCooldown = moveDelay
	}

	return nil
}

func draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	for i := range gameMap {
		for j := range gameMap[i] {
			var c color.Color
			if gameMap[i][j] == 0 {
				c = color.RGBA{255, 255, 255, 255}
			} else {
				c = color.RGBA{0, 0, 255, 255}
			}
			vector.DrawFilledRect(screen, float32(j*tileSize), float32(i*tileSize), tileSize, tileSize, c, false)
		}
	}
	vector.DrawFilledRect(screen, float32(playerX*tileSize), float32(playerY*tileSize), tileSize, tileSize, color.RGBA{255, 0, 0, 255}, false)
}

func main() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return update(nil)
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
