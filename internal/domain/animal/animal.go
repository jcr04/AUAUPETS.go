package animal

type Animal struct {
	ID    string
	Name  string
	Breed string
	Age   string
}

func NewAnimal(id, name, breed, age string) *Animal {
	return &Animal{
		ID:    id,
		Name:  name,
		Breed: breed,
		Age:   age,
	}
}
