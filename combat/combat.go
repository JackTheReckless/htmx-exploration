package combat

type CombatDTO struct {
	ID     int
	Name   string
	Target string
	Type   string
	Amt    int
}

type User interface {
	GetName() string
	GetHealth() int
	GetMaxHealth() int
	TakeDamage(int)
	Heal(int)
	IncreaseEXP(int)
}

type Enemy interface {
	GetName() string
	GetHealth() int
	GetMaxHealth() int
	TakeDamage(int)
	Heal(int)
}
