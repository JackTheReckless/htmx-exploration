package forest

import (
	"htmx-exploration/combat"
	"htmx-exploration/enemy"
	"math/rand"
	"sync"
	"time"
)

type ForestService struct {
}

func NewForestService() *ForestService {
	return &ForestService{}
}

var forestMutex sync.Mutex
var currentForest *Forest

func NewForest() *Forest {
	forest := &Forest{
		EncounterMap: make(map[string]EncounterFunction),
	}

	forest.RegisterEncounter("basicCombat", forest.EncounterBasicEnemy)

	return forest
}

func (f *Forest) RegisterEncounter(encounterType string, encounter EncounterFunction) {
	f.EncounterMap[encounterType] = encounter
}

func SetCurrentForest(forest *Forest) {
	forestMutex.Lock()
	defer forestMutex.Unlock()
	currentForest = forest
}

func GetCurrentForest() *Forest {
	forestMutex.Lock()
	defer forestMutex.Unlock()
	return currentForest
}

func (f *Forest) RandomEncounter() {
	// Seed the random number generator
	rand.Intn(int(time.Now().UnixNano()))

	// Get a random encounter ID
	encounterTypes := make([]string, 0, len(f.EncounterMap))
	for eType := range f.EncounterMap {
		encounterTypes = append(encounterTypes, eType)
	}
	randomType := encounterTypes[rand.Intn(len(encounterTypes))]

	// Execute the random encounter
	if encounter, ok := f.EncounterMap[randomType]; ok {
		f.currentEncounter = randomType
		encounter()
	}
}

func (f *Forest) EncounterBasicEnemy() {
	e := enemy.NewEnemy()

	enemy.SetCurrentEnemy(e)
	combat.InitiateCombat()
}

func ResetForest() {
	forestMutex.Lock()
	defer forestMutex.Unlock()

	currentForest = nil
}
