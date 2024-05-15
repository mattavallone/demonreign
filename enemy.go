package main

import (
	"math"
	"math/rand"
	"slices"
	"time"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

var (
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	numEnemies = random.Intn(5) + 1 // add 1 to guarantee at least one enemy spawns
	direction  = []string{"N", "S", "W", "E"}
)

type gameEnemy struct {
	x, y        int
	moveDelay   int
	orientation string
	health      int

	animInstance Animation
	fly          Animation
	attack       Animation
	takeHit      Animation
	death        Animation
}

func LoadEnemies() []*gameEnemy {
	gameEnemies := make([]*gameEnemy, numEnemies)
	for i := 0; i < numEnemies; i++ {
		gameEnemies[i] = &gameEnemy{
			x:           random.Intn(mapSize),
			y:           random.Intn(mapSize),
			health:      4,
			moveDelay:   10,           // Adjust this value to change the speed of the enemy
			orientation: direction[1], // S
		}
		gameEnemies[i].generateSpawnPosition()
		gameEnemies[i].LoadEnemyImage()
	}
	return gameEnemies
}

func (nme *gameEnemy) generateSpawnPosition() {
	for gameMap[nme.y][nme.x] != 0 { // don't spawn on other objects
		nme.x = random.Intn(mapSize)
		nme.y = random.Intn(mapSize)
		nme.orientation = direction[random.Intn(4)]
	}
	gameMap[nme.y][nme.x] = 2 // update the map to indicate where the enemy spawns
}

func (nme *gameEnemy) LoadEnemyImage() {
	attackImg, _, _ := ebitenutil.NewImageFromFile("assets/Monsters_Creatures_Fantasy/Flying eye/Attack.png")
	flyImg, _, _ := ebitenutil.NewImageFromFile("assets/Monsters_Creatures_Fantasy/Flying eye/Flight.png")
	deathImg, _, _ := ebitenutil.NewImageFromFile("assets/Monsters_Creatures_Fantasy/Flying eye/Death.png")
	takeHitImg, _, _ := ebitenutil.NewImageFromFile("assets/Monsters_Creatures_Fantasy/Flying eye/Take Hit.png")

	gAttack := ganim8.NewGrid(145, 150, attackImg.Bounds().Dx(), attackImg.Bounds().Dy(), 62, 60, 1)
	gFly := ganim8.NewGrid(145, 150, flyImg.Bounds().Dx(), flyImg.Bounds().Dy(), 62, 60, 1)
	gDeath := ganim8.NewGrid(145, 150, deathImg.Bounds().Dx(), deathImg.Bounds().Dy(), 62, 60, 1)
	gTakeHit := ganim8.NewGrid(145, 150, takeHitImg.Bounds().Dx(), takeHitImg.Bounds().Dy(), 62, 60, 1)

	nme.fly.anim = ganim8.New(flyImg, gFly.Frames("1-8", "1"), 10*time.Millisecond)
	nme.fly.originX = 0
	nme.fly.originY = 0

	nme.attack.anim = ganim8.New(attackImg, gAttack.Frames("1-8", "1"), 10*time.Millisecond)
	nme.attack.originX = 0
	nme.attack.originY = 0

	nme.death.anim = ganim8.New(deathImg, gDeath.Frames("1-4", "1"), 10*time.Millisecond)
	nme.death.originX = 0
	nme.death.originY = 0

	nme.takeHit.anim = ganim8.New(takeHitImg, gTakeHit.Frames("1-4", "1"), 10*time.Millisecond)
	nme.takeHit.originX = 0
	nme.takeHit.originY = 0

	nme.animInstance = nme.fly
}

func (nme *gameEnemy) ChangeDir() {
	i := slices.Index(direction, nme.orientation) + 1
	if i == 4 {
		i = 0
	}
	newDir := direction[i]
	nme.orientation = newDir
}

func (nme *gameEnemy) Move() {
	switch nme.orientation {
	case "E":
		nme.x = int(math.Min(float64(nme.x+1), mapSize-1))
	case "W":
		nme.x = int(math.Max(float64(nme.x-1), 0))
	case "N":
		nme.y = int(math.Max(float64(nme.y-1), 0))
	case "S":
		nme.y = int(math.Min(float64(nme.y+1), mapSize-1))
	}
}
