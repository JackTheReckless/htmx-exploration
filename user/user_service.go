package user

import (
	"strings"
	"sync"
)

type UserService struct {
}

var userMutex sync.Mutex
var currentUser *User
var gmNames = []string{"jon", "gm", "admin", "jack"}

func NewUserService() *UserService {
	return &UserService{}
}

func isUserGameMaster(name string) bool {
	nameLower := strings.ToLower(name)
	for _, gmName := range gmNames {
		if nameLower == gmName {
			return true
		}
	}
	return false
}

func (us *UserService) AwardExp(user *User, amount int) {
	user.IncreaseEXP(amount)
}

func NewUser(name string, class string) *User {
	isGameMaster := isUserGameMaster(name)

	return &User{
		Name:          name,
		Class:         class,
		Level:         1,
		IsGameMaster:  isGameMaster,
		Health:        10,
		MaxHealth:     10,
		PercentHealth: 100,
		Stamina:       10,
		MaxStamina:    10,
		Title:         "Initiate",
		EXP:           0,
		AwardedEXP:    0,
		Strength:      10,
		Dexterity:     10,
		Constitution:  10,
		Intelligence:  10,
		Wisdom:        10,
		Charisma:      10,
		Luck:          10,
	}
}

// SetCurrentUser sets the global user instance
func SetCurrentUser(user *User) {
	userMutex.Lock()
	defer userMutex.Unlock()
	currentUser = user
}

// GetCurrentUser retrieves the global user instance
func GetCurrentUser() *User {
	userMutex.Lock()
	defer userMutex.Unlock()
	return currentUser
}

// GetName returns the name of the user
func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetHealth() int {
	return u.Health
}

func (u *User) GetMaxHealth() int {
	return u.MaxHealth
}

// TakeDamage deducts the specified amount from the user's health
func (u *User) TakeDamage(amount int) {
	u.Health -= amount
	if u.Health < 0 {
		u.Health = 0
	}
	u.PercentHealth = int(float32(u.Health) / float32(u.MaxHealth) * 100)
}

func (u *User) SetClass(class string) {
	u.Class = class
}

func (u *User) Heal(amount int) {
	u.Health += amount
	if u.Health > u.MaxHealth {
		u.Health = u.MaxHealth
	}

	u.PercentHealth = int(float32(u.Health) / float32(u.MaxHealth) * 100)
}

func (u *User) SetCurrentHealth(health int) {
	u.Health = health
	u.PercentHealth = int(float32(u.Health) / float32(u.MaxHealth) * 100)
}

func (u *User) SetCurrentStamina(stamina int) {
	u.Stamina = stamina
}

func (u *User) IncreaseEXP(amount int) {
	u.AwardedEXP = amount
	u.EXP += amount
}

// DeleteUser deletes the current user
func DeleteUser() {
	userMutex.Lock()
	defer userMutex.Unlock()

	currentUser = nil
}
