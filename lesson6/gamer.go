package main

import "unsafe"

type (
	firstBitPositionType uint8
	bitsCountType        uint8
)

const (
	goldStart     firstBitPositionType = 0
	goldBitsCount bitsCountType        = 31

	manaStart     firstBitPositionType = 0
	manaBitsCount bitsCountType        = 10

	healthStart     firstBitPositionType = 0
	healthBitsCount bitsCountType        = 10

	respectStart     firstBitPositionType = 10
	respectBitsCount bitsCountType        = 4

	strengthStart     firstBitPositionType = 10
	strengthBitsCount bitsCountType        = 4

	experienceStart     firstBitPositionType = 4
	experienceBitsCount bitsCountType        = 4

	levelStart     firstBitPositionType = 0
	levelBitsCount bitsCountType        = 4

	houseStart     firstBitPositionType = 14
	houseBitsCount bitsCountType        = 1

	weaponStart     firstBitPositionType = 15
	weaponBitsCount bitsCountType        = 1

	familyStart     firstBitPositionType = 31
	familyBitsCount bitsCountType        = 1

	typeStart     firstBitPositionType = 14
	typeBitsCount bitsCountType        = 2
)

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

		person.UserName.NameLen = uint8(n)
		copy(person.UserName.UserName[:n], name[:n])
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.CoorX = int32(x)
		person.CoorY = int32(y)
		person.CoorZ = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.FamilyGold = setNewData(person.FamilyGold, uint32(gold), goldStart, goldBitsCount)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.WeaponHomeRespectMana = setNewData(person.WeaponHomeRespectMana, uint16(mana), manaStart, manaBitsCount)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.TypeStrengthHealth = setNewData(person.TypeStrengthHealth, uint16(health), healthStart, healthBitsCount)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.WeaponHomeRespectMana = setNewData(person.WeaponHomeRespectMana, uint16(respect), respectStart, respectBitsCount)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.TypeStrengthHealth = setNewData(person.TypeStrengthHealth, uint16(strength), strengthStart, strengthBitsCount)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.ExperienceLevel = setNewData(person.ExperienceLevel, uint8(experience), experienceStart, experienceBitsCount)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.ExperienceLevel = setNewData(person.ExperienceLevel, uint8(level), levelStart, levelBitsCount)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.WeaponHomeRespectMana = setNewData(person.WeaponHomeRespectMana, 1, houseStart, houseBitsCount)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.WeaponHomeRespectMana = setNewData(person.WeaponHomeRespectMana, 1, weaponStart, weaponBitsCount)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.FamilyGold = setNewData(person.FamilyGold, 1, familyStart, familyBitsCount)
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.TypeStrengthHealth = setNewData(person.TypeStrengthHealth, uint16(personType), typeStart, typeBitsCount)
	}
}

func setNewData[T uint8 | uint32 | uint16](data, newData T, start firstBitPositionType, bitsCount bitsCountType) T {
	return resetBits(data, start, bitsCount) | (newData << start)
}

func resetBits[T uint8 | uint32 | uint16](data T, start firstBitPositionType, dataBitsCount bitsCountType) T {
	genericBitsCount := unsafe.Sizeof(T(0)) * 8
	resultMask := T(1<<genericBitsCount - 1)
	xorMask := T((1<<dataBitsCount - 1) << start)
	resultMask ^= xorMask

	return data & resultMask
}

func getData[T uint8 | uint32 | uint16](data T, start firstBitPositionType, bitsCount bitsCountType) int {
	mask := T(1<<bitsCount - 1)
	return int((data >> start) & mask)
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type userNameType struct {
	UserName [42]byte `json:"username"`
	NameLen  uint8    `json:"name"`
}

type GamePerson struct {
	UserName        userNameType `json:"user_name"`
	ExperienceLevel uint8        `json:"experience_level"`

	CoorX int32 `json:"coor_x"`
	CoorY int32 `json:"coor_y"`
	CoorZ int32 `json:"coor_z"`

	// make serialization for family and gold separately
	FamilyGold uint32 `json:"family_gold"`

	WeaponHomeRespectMana uint16 `json:"weapon_home_respect_mana"`
	TypeStrengthHealth    uint16 `json:"type_strength_health"`
}

func NewGamePerson(options ...Option) GamePerson {
	gamer := GamePerson{}

	for _, option := range options {
		option(&gamer)
	}

	return gamer
}

func (p *GamePerson) Name() string {
	return string(p.UserName.UserName[:p.UserName.NameLen])
}

func (p *GamePerson) X() int {
	return int(p.CoorX)
}

func (p *GamePerson) Y() int {
	return int(p.CoorY)
}

func (p *GamePerson) Z() int {
	return int(p.CoorZ)
}

func (p *GamePerson) Gold() int {
	return getData(p.FamilyGold, goldStart, goldBitsCount)
}

func (p *GamePerson) Mana() int {
	return getData(p.WeaponHomeRespectMana, manaStart, manaBitsCount)
}

func (p *GamePerson) Health() int {
	return getData(p.TypeStrengthHealth, healthStart, healthBitsCount)
}

func (p *GamePerson) Respect() int {
	return getData(p.WeaponHomeRespectMana, respectStart, respectBitsCount)
}

func (p *GamePerson) Strength() int {
	return getData(p.TypeStrengthHealth, strengthStart, strengthBitsCount)
}

func (p *GamePerson) Experience() int {
	return getData(p.ExperienceLevel, experienceStart, experienceBitsCount)
}

func (p *GamePerson) Level() int {
	return getData(p.ExperienceLevel, levelStart, levelBitsCount)
}

func (p *GamePerson) HasHouse() bool {
	return getData(p.WeaponHomeRespectMana, houseStart, houseBitsCount) == 1
}

func (p *GamePerson) HasGun() bool {
	return getData(p.WeaponHomeRespectMana, weaponStart, weaponBitsCount) == 1
}

func (p *GamePerson) HasFamilty() bool {
	return getData(p.FamilyGold, familyStart, familyBitsCount) == 1
}

func (p *GamePerson) Type() int {
	return getData(p.TypeStrengthHealth, typeStart, typeBitsCount)
}
