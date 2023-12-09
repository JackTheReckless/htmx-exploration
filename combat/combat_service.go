package combat

import (
	"math/rand"
	"sort"
	"sync"
)

type CombatService struct {
}

func NewCombatService() *CombatService {
	return &CombatService{}
}

var (
	Combat        []CombatDTO
	combatCounter int
	combatMu      sync.RWMutex
)

func InitiateCombat() {
	combatCounter = 1
}

func NewCombatRound() {
	combatCounter += 1
}

func UserAttack(user User, enemy Enemy) {
	combatMu.Lock()
	defer combatMu.Unlock()

	damage := calcDamage()

	combatCounter += 1
	userAttack := CombatDTO{
		ID:     combatCounter,
		Name:   "You",
		Target: enemy.GetName(),
		Type:   "dmg",
		Amt:    damage,
	}
	enemy.TakeDamage(damage)

	Combat = append(Combat, userAttack)

	if enemy.GetHealth() > 0 {
		EnemyAttack(user, enemy)
	} else {
		expAmount := calculateExpAmount()
		user.IncreaseEXP(expAmount)
	}
}

func calculateExpAmount() int {
	return rand.Intn(100-1) + 1 // equivalent to rand.Range(min, max int)
}

func EnemyAttack(user User, enemy Enemy) {
	damage := calcDamage()

	combatCounter += 1
	enemyAttack := CombatDTO{
		ID:     combatCounter,
		Name:   enemy.GetName(),
		Target: "You",
		Type:   "dmg",
		Amt:    damage,
	}
	user.TakeDamage(damage)

	Combat = append(Combat, enemyAttack)
}

func UserHeal(user User) {

	if user.GetHealth() < user.GetMaxHealth() {

		healing := calcHealing()

		combatCounter += 1
		userHealing := CombatDTO{
			ID:     combatCounter,
			Name:   user.GetName(),
			Target: user.GetName(),
			Type:   "heal",
			Amt:    healing,
		}
		user.Heal(healing)

		Combat = append(Combat, userHealing)
	}
}

func EnemyHeal(enemy Enemy) {

	if enemy.GetHealth() < enemy.GetMaxHealth() {
		healing := calcHealing()

		combatCounter += 1
		enemyHealing := CombatDTO{

			ID:     combatCounter,
			Name:   enemy.GetName(),
			Target: enemy.GetName(),
			Type:   "heal",
			Amt:    healing,
		}
		enemy.Heal(healing)

		Combat = append(Combat, enemyHealing)
	}
}

func calcHealing() int {
	return rand.Intn(5) + 1 // Adding 1 to avoid zero healing for now
}

func calcDamage() int {
	return rand.Intn(5) + 1 // Adding 1 to avoid zero damage for now
}

// GetCombatLog returns the current attacks array
// sorted descending so newest is at the top
func GetCombatLog() []CombatDTO {
	combatMu.Lock()
	defer combatMu.Unlock()

	reversedAttacks := make([]CombatDTO, len(Combat))
	copy(reversedAttacks, Combat)

	sort.Slice(reversedAttacks, func(i, j int) bool {
		return reversedAttacks[i].ID > reversedAttacks[j].ID
	})

	return reversedAttacks
}

func ResetCombat() {
	combatMu.Lock()
	defer combatMu.Unlock()

	Combat = nil
}
