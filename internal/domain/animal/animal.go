package animal

type Animal struct {
	ID    string
	Name  string
	Breed string
}

func NewAnimal(id, name, breed string) *Animal {
	return &Animal{id, name, breed}
}
