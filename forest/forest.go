package forest

type EncounterFunction func()

type Forest struct {
	EncounterMap     map[string]EncounterFunction
	currentEncounter string
}
