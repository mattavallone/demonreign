package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	tileSize     = 16
	mapSize      = 16
)

var (
	moveCooldown = 0

	floorsImage *ebiten.Image
	wallsImage  *ebiten.Image

	gameMap = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
)

func init() {
	var err error
	floorsImage, _, err = ebitenutil.NewImageFromFile("assets/atlas_floor-16x16.png")
	if err != nil {
		log.Fatal(err)
	}
	wallsImage, _, err = ebitenutil.NewImageFromFile("assets/atlas_walls_low-16x16.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Demon Reign")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := &Game{
		player: NewPlayer(),
	}

	LoadPlayerImage(g.player)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player *gamePlayer
}

func (g *Game) Update() error {
	if moveCooldown > 0 {
		moveCooldown--
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.player.y > 0 && gameMap[g.player.y-1][g.player.x] != 1 {
			g.player.y--
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "N"
		g.player.animInstance = g.player.walkup
		g.player.animInstance.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.player.y < mapSize-1 && gameMap[g.player.y+1][g.player.x] != 1 {
			g.player.y++
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "S"

		g.player.animInstance = g.player.walkdown
		g.player.animInstance.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if g.player.x > 0 && gameMap[g.player.y][g.player.x-1] != 1 {
			g.player.x--
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "W"

		g.player.animInstance = g.player.walkleft
		g.player.animInstance.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.player.x < mapSize-1 && gameMap[g.player.y][g.player.x+1] != 1 {
			g.player.x++
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "E"

		g.player.animInstance = g.player.walkright
		g.player.animInstance.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		switch g.player.orientation {
		case "E":
			g.player.animInstance = g.player.meleeright
		case "W":
			g.player.animInstance = g.player.meleeleft
		case "N":
			g.player.animInstance = g.player.meleeup
		case "S":
			g.player.animInstance = g.player.meleedown
		default:
			fmt.Printf("game.Update() - orientation not recognized")
			g.player.animInstance = g.player.meleeright
		}

		g.player.animInstance.Update()

	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMap(screen)
	g.drawPlayer(screen)
}

func (g *Game) drawMap(screen *ebiten.Image) {
	for y := range gameMap {
		for x := range gameMap[y] {
			// floors
			tileIndex := gameMap[y][x]

			imageToRender := floorsImage

			srcRect := image.Rect(0, 0, 16, 16)
			dstRect := image.Rect(x*16, y*16, (x+1)*16, (y+1)*16)

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(dstRect.Min.X), float64(dstRect.Min.Y))
			screen.DrawImage(imageToRender.SubImage(srcRect).(*ebiten.Image), opts)

			// walls
			if tileIndex == 1 {
				srcRect := image.Rect(0, 0, 16, 16)
				dstRect := image.Rect(x*16, y*16, (x+1)*16, (y+1)*16)

				if (gameMap[y][x+1] == 1) && (gameMap[y+1][x] == 1) { // top left corner
					srcRect = image.Rect(16, 0, 32, 16)
				} else if (gameMap[y][x-1] == 1) && (gameMap[y+1][x] == 1) { // top right corner
					srcRect = image.Rect(48, 0, 64, 16)
				} else if (gameMap[y-1][x] == 1) && (gameMap[y][x+1] == 1) { // bottom left corner
					srcRect = image.Rect(16, 32, 32, 48)
				} else if (gameMap[y-1][x] == 1) && (gameMap[y][x-1] == 1) { //bottom right corner
					srcRect = image.Rect(48, 48, 64, 64)
				} else if (gameMap[y-1][x] == 1) && (gameMap[y+1][x] == 1) || (gameMap[y+1][x] == 1) && (gameMap[y-1][x] == 0) || (gameMap[y+1][x] == 0) && (gameMap[y-1][x] == 1) { // vertical side
					srcRect = image.Rect(0, 0, 16, 16)
				} else if (gameMap[y][x+1] == 1) && (gameMap[y][x-1] == 1) || (gameMap[y][x+1] == 1) && (gameMap[y][x-1] == 0) || (gameMap[y][x+1] == 0) && (gameMap[y][x-1] == 1) { // hortizontal side
					srcRect = image.Rect(24, 48, 40, 64)
				}
				imageToRender := wallsImage

				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(dstRect.Min.X), float64(dstRect.Min.Y))
				screen.DrawImage(imageToRender.SubImage(srcRect).(*ebiten.Image), opts)
			}
		}
	}
}

func (g *Game) drawPlayer(screen *ebiten.Image) {
	dstRect := image.Rect(g.player.x*tileSize, g.player.y*tileSize, (g.player.x+1)*tileSize, (g.player.y+1)*tileSize)

	g.player.animInstance.Draw(screen, ganim8.DrawOpts(float64(dstRect.Min.X), float64(dstRect.Min.Y), 0, 1, 1, 0.2, 0.2))
}
