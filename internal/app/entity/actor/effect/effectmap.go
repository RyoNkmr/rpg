package effect

type EffectMap map[Effect]bool

func (m *EffectMap) Clear() {
	*m = EffectMap{}
}

func (m EffectMap) Has(e Effect) (has bool) {
	v, ok := m[e]
	has = ok && v
	return has
}

func (m EffectMap) Add(e Effect) {
	m[e] = true
}

func (m EffectMap) Remove(e Effect) {
	m[e] = false
}

func (m EffectMap) AsOrderedList() (list []Effect) {
	for _, e := range EffectList {
		if m.Has(e) {
			list = append(list, e)
		}
	}
	return list
}
