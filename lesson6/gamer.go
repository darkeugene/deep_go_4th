package main

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		n := len(name)
		if n == 0 {
			return
		}
		if n > 42 {
			n = 42
		}

		person.nameLen = uint8(n)
		copy(person.userName[:n], name[:n])
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.familyGold = (person.familyGold & 1 << 31) | uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.weaponHomeRespectMana = (person.weaponHomeRespectMana & (0b111111 << 10)) | uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.typeStrengthHealth = (person.typeStrengthHealth & (0b111111 << 10)) | uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.weaponHomeRespectMana = (person.weaponHomeRespectMana & (0b11<<14 | (0b1<<10 - 1))) | uint16(respect)<<10
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.typeStrengthHealth = (person.typeStrengthHealth & (0b11<<14 | (0b1<<10 - 1))) | uint16(strength)<<10
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceLevel = (person.experienceLevel & 0b1111) | uint8(experience)<<4
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.experienceLevel = (person.experienceLevel & (0b1111 << 4)) | uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.weaponHomeRespectMana = (person.weaponHomeRespectMana & (0b1<<15 | (0b1<<14 - 1))) | 1<<14
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.weaponHomeRespectMana = (person.weaponHomeRespectMana & (0b1<<15 - 1)) | 1<<15
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.familyGold = (person.familyGold & (1<<31 - 1)) | 1<<31
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.typeStrengthHealth = (person.typeStrengthHealth & (0b1<<14 - 1)) | uint16(personType)<<14
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	userName [42]byte

	nameLen         uint8
	experienceLevel uint8

	x int32
	y int32
	z int32

	familyGold uint32

	weaponHomeRespectMana uint16
	typeStrengthHealth    uint16
}

func NewGamePerson(options ...Option) GamePerson {
	gamer := GamePerson{}

	for _, option := range options {
		option(&gamer)
	}

	return gamer
}

func (p *GamePerson) Name() string {
	// need to implement
	return string(p.userName[:p.nameLen])
}

func (p *GamePerson) X() int {
	// need to implement
	return int(p.x)
}

func (p *GamePerson) Y() int {
	// need to implement
	return int(p.y)
}

func (p *GamePerson) Z() int {
	// need to implement
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	// need to implement
	return int(p.familyGold & (1<<31 - 1))
}

func (p *GamePerson) Mana() int {
	// need to implement
	return int(p.weaponHomeRespectMana & (1<<10 - 1))
}

func (p *GamePerson) Health() int {
	// need to implement
	return int(p.typeStrengthHealth & (1<<10 - 1))
}

func (p *GamePerson) Respect() int {
	// need to implement
	return int(p.weaponHomeRespectMana >> 10 & 0b1111)
}

func (p *GamePerson) Strength() int {
	// need to implement
	return int(p.typeStrengthHealth >> 10 & 0b1111)
}

func (p *GamePerson) Experience() int {
	// need to implement
	return int(p.experienceLevel >> 4 & 0b1111)
}

func (p *GamePerson) Level() int {
	// need to implement
	return int(p.experienceLevel & 0b1111)
}

func (p *GamePerson) HasHouse() bool {
	// need to implement
	return (p.weaponHomeRespectMana >> 14 & 1) == 1
}

func (p *GamePerson) HasGun() bool {
	// need to implement
	return (p.weaponHomeRespectMana >> 15 & 1) == 1
}

func (p *GamePerson) HasFamilty() bool {
	// need to implement
	return (p.familyGold >> 31 & 1) == 1
}

func (p *GamePerson) Type() int {
	// need to implement
	return int(p.typeStrengthHealth >> 14 & 0b11)
}
