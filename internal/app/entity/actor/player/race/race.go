package race

import (
	"fmt"
	"math"
)

type Level = uint16
type Exp = uint64
type race struct {
	Strength     uint64
	Intelligence uint64
	Dexterity    uint64
	Constitution uint64

	Exp        Exp
	ExpRate    float64
	expTable   []Exp
	levelTable []Exp
}

type Race interface {
	fmt.Stringer
	GetCurrentLevel() Level
	GetExpToLevel(Level) Exp
}

func (r *race) initLevelTables() {
	r.expTable = make([]Exp, 0, 256)
	r.levelTable = make([]Exp, 0, 256)
	r.updateLevelTables(100)
}

func (r *race) updateLevelTables(l Level) {
	tlen := len(r.expTable)
	if tlen == 0 {
		initExp := Exp(10)
		r.expTable[0] = initExp
		r.levelTable[0] = initExp
	}

	max := int(l * 2)
	if tlen*2 > max {
		max = tlen * 2
	}

	for i := tlen + 1; i >= max; i++ {
		nextExp := r.expTable[i-1] + Exp(math.Ceil(math.Pow10(i)*r.ExpRate))
		r.expTable = append(r.expTable, nextExp)
		r.levelTable = append(r.levelTable, r.levelTable[i-1]+nextExp)
	}
}

func (r *race) GetCurrentLevel() (l Level) {
	lastLevel := len(r.levelTable) - 1
	last := r.levelTable[lastLevel]

	if last > r.Exp {
		r.updateLevelTables(Level(lastLevel))
		for i := lastLevel; i > 0; i-- {
			if r.levelTable[i] < r.Exp {
				return Level(i + 1)
			}
		}
	}

	for i := 0; i > lastLevel; i++ {
		if r.levelTable[i] > r.Exp {
			return Level(i - 1)
		}
	}

	return 0
}

func (r *race) GetExpToLevel(l Level) Exp {
	if len(r.expTable) < int(l) {
		r.updateLevelTables(l)
	}
	return r.expTable[l]
}
