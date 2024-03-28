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
)

var (
	moveCooldown = 0

	iconImage   *ebiten.Image
	floorsImage *ebiten.Image
	wallsImage  *ebiten.Image
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

	img, _, _ := ebitenutil.NewImageFromFile("assets/characters/link/link.png")
	crop := image.Rect(28, 96, 44, 115)
	iconImage = img.SubImage(crop).(*ebiten.Image)

}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Demon Reign")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowIcon([]image.Image{iconImage})

	g := &Game{
		player:  NewPlayer(),
		enemies: LoadEnemies(),
	}

	LoadPlayerImage(g.player)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player  *gamePlayer
	enemies []*gameEnemy
}

func (g *Game) Update() error {
	if moveCooldown > 0 {
		moveCooldown--
		return nil
	}

	for _, nme := range g.enemies {
		moveCooldown = nme.moveDelay
		nme.animInstance.anim.Update()
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.player.y > 0 && gameMap[g.player.y-1][g.player.x] < 1 {
			g.player.y--
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "N"
		g.player.animInstance = g.player.walkup
		g.player.animInstance.anim.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.player.y < mapSize-1 && gameMap[g.player.y+1][g.player.x] < 1 {
			g.player.y++
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "S"

		g.player.animInstance = g.player.walkdown
		g.player.animInstance.anim.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if g.player.x > 0 && gameMap[g.player.y][g.player.x-1] < 1 {
			g.player.x--
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "W"

		g.player.animInstance = g.player.walkleft
		g.player.animInstance.anim.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.player.x < mapSize-1 && gameMap[g.player.y][g.player.x+1] < 1 {
			g.player.x++
		}
		moveCooldown = g.player.moveDelay
		g.player.orientation = "E"

		g.player.animInstance = g.player.walkright
		g.player.animInstance.anim.Update()

	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		switch g.player.orientation {
		case "E":
			g.player.animInstance = g.player.meleeright
			if gameMap[g.player.y][g.player.x+1] == 2 {
				g.enemyTakeHit(g.player.x+1, g.player.y)
			}
		case "W":
			g.player.animInstance = g.player.meleeleft
			if gameMap[g.player.y][g.player.x-1] == 2 {
				g.enemyTakeHit(g.player.x-1, g.player.y)
			}
		case "N":
			g.player.animInstance = g.player.meleeup
			if gameMap[g.player.y-1][g.player.x] == 2 {
				g.enemyTakeHit(g.player.x, g.player.y-1)
			}
		case "S":
			g.player.animInstance = g.player.meleedown
			if gameMap[g.player.y+1][g.player.x] == 2 {
				g.enemyTakeHit(g.player.x, g.player.y+1)
			}
		default:
			fmt.Printf("game.Update() - orientation not recognized")
			g.player.animInstance = g.player.meleeright
		}

		g.player.animInstance.anim.Update()

	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMap(screen)
	g.drawPlayer(screen)
	g.drawEnemies(screen)
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
				} else if (gameMap[y-1][x] == 1) && (gameMap[y+1][x] == 1) || (gameMap[y+1][x] == 1) && (gameMap[y-1][x] == 0) { // vertical side
					srcRect = image.Rect(0, 0, 16, 16)
				} else if (gameMap[y+1][x] == 0) && (gameMap[y-1][x] == 1) { // vertical bottom end
					srcRect = image.Rect(0, 32, 16, 48)
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

	g.player.animInstance.anim.Draw(screen, ganim8.DrawOpts(float64(dstRect.Min.X), float64(dstRect.Min.Y), 0, 1, 1, g.player.animInstance.originX, g.player.animInstance.originY))
}

func (g *Game) drawEnemies(screen *ebiten.Image) {
	scaleX := screenWidth / float64(mapSize*tileSize)
	scaleY := screenHeight / float64(mapSize*tileSize)
	for i := 0; i < len(g.enemies); i++ {
		dstRect := image.Rect(g.enemies[i].x*tileSize, g.enemies[i].y*tileSize, (g.enemies[i].x+1)*tileSize, (g.enemies[i].y+1)*tileSize)

		g.enemies[i].animInstance.anim.Draw(screen, ganim8.DrawOpts(float64(dstRect.Min.X), float64(dstRect.Min.Y), 0, 1/scaleX, 1/scaleY, g.enemies[i].animInstance.originX, g.enemies[i].animInstance.originY))
	}

}

func (g *Game) enemyTakeHit(x, y int) {
	i := len(g.enemies) - 1
	for len(g.enemies) > 0 {
		nme := g.enemies[i]
		fmt.Println(i)
		if nme.x == x && nme.y == y {
			nme.health--
			if nme.health > 0 { // enemy still alive
				nme.animInstance = nme.takeHit
				nme.animInstance.anim.Update()
			} else { // enemy is dead
				nme.animInstance = nme.death
				gameMap[y][x] = 0
				nme.animInstance.anim.Update()
				g.enemies[i] = g.enemies[len(g.enemies)-1]
				g.enemies = g.enemies[:len(g.enemies)-1]
			}
			break
		}
		i--
	}
}
