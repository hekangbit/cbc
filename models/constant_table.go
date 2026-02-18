package models

type ConstantTable struct {
	entries map[string]*ConstantEntry
	order   []string
}

func NewConstantTable() *ConstantTable {
	return &ConstantTable{
		entries: make(map[string]*ConstantEntry),
		order:   []string{},
	}
}

func (this *ConstantTable) IsEmpty() bool {
	return len(this.entries) == 0
}

func (this *ConstantTable) Intern(s string) *ConstantEntry {
	if ent, ok := this.entries[s]; ok {
		return ent
	}
	ent := NewConstantEntry(s)
	this.entries[s] = ent
	this.order = append(this.order, s)
	return ent
}

func (this *ConstantTable) Entries() []*ConstantEntry {
	result := make([]*ConstantEntry, 0, len(this.order))
	for _, key := range this.order {
		result = append(result, this.entries[key])
	}
	return result
}

// TODO: how to call
func (this *ConstantTable) ForEach(f func(*ConstantEntry)) {
	for _, key := range this.order {
		f(this.entries[key])
	}
}
