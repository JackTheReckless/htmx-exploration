package enemy

import (
	"math/rand"
	"sync"
)

type EnemyService struct {
}

func NewEnemyService() *EnemyService {
	return &EnemyService{}
}

var enemyMutex sync.Mutex
var currentEnemy *Enemy

func SetCurrentEnemy(enemy *Enemy) {
	enemyMutex.Lock()
	defer enemyMutex.Unlock()
	currentEnemy = enemy
}

func GetCurrentEnemy() *Enemy {
	enemyMutex.Lock()
	defer enemyMutex.Unlock()
	return currentEnemy
}

// GetName returns the name of the user
func (e *Enemy) GetName() string {
	return e.Name
}

func (e *Enemy) GetHealth() int {
	return e.Health
}

func (e *Enemy) GetMaxHealth() int {
	return e.MaxHealth
}

// TakeDamage deducts the specified amount from the user's health
func (e *Enemy) TakeDamage(amount int) {
	e.Health -= amount
	if e.Health < 0 {
		e.Health = 0
	}
	e.PercentHealth = int(float32(e.Health) / float32(e.MaxHealth) * 100)
}

func (e *Enemy) Heal(amount int) {
	e.Health += amount
	if e.Health > e.MaxHealth {
		e.Health = e.MaxHealth
	}
	e.PercentHealth = int(float32(e.Health) / float32(e.MaxHealth) * 100)
}

func (e *Enemy) SetCurrentHealth(health int) {
	e.Health += health
	if e.Health > e.MaxHealth {
		e.Health = e.MaxHealth
	}

	e.PercentHealth = int(float32(e.Health) / float32(e.MaxHealth) * 100)
}

func NewEnemy() *Enemy {
	return generateRandomEnemy()
}

func generateRandomEnemy() *Enemy {
	enemies := []Enemy{
		{
			Name:          "Innocent-Looking Squirrel",
			Health:        10,
			MaxHealth:     10,
			PercentHealth: 100,
			Image:         "../static/Innocent-Looking Squirrel Lv1.gif",
			Weapon:        "Nuts",
		},
	}
	return &enemies[rand.Intn(len(enemies))]
}

// DeleteUser deletes the current user
func ResetEnemy() {
	enemyMutex.Lock()
	defer enemyMutex.Unlock()

	currentEnemy = nil
}
