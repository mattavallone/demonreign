package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

var (
	random     = rand.New(rand.NewSource(time.Now().UnixNano()))
	numEnemies = random.Intn(10) + 1 // add 1 to guarantee at least one enemy spawns
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
			moveDelay:   10,           // Adjust this value to change the speed of the player
			orientation: direction[1], // S
		}
		gameEnemies[i].generateSpawnPosition()
		gameEnemies[i].LoadEnemyImage()
	}
	return gameEnemies
}

func (nme *gameEnemy) generateSpawnPosition() {
	for gameMap[nme.y][nme.x] == 1 {
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

	gAttack := ganim8.NewGrid(150, 150, attackImg.Bounds().Dx(), attackImg.Bounds().Dy(), 56, 60, 1)
	gFly := ganim8.NewGrid(150, 150, flyImg.Bounds().Dx(), flyImg.Bounds().Dy(), 56, 60, 1)
	gDeath := ganim8.NewGrid(150, 150, deathImg.Bounds().Dx(), deathImg.Bounds().Dy(), 56, 60, 1)
	gTakeHit := ganim8.NewGrid(150, 150, takeHitImg.Bounds().Dx(), takeHitImg.Bounds().Dy(), 56, 60, 1)

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
