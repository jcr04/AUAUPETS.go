package animal

type Animal struct {
	ID       string
	Name     string
	Breed    string
	Age      int
	CheckIn  string // formato YYYY-MM-DD
	CheckOut string // formato YYYY-MM-DD
}

func NewAnimal(id, name, breed string, age int, CheckIn, CheckOut string) *Animal {
	return &Animal{
		ID:       id,
		Name:     name,
		Breed:    breed,
		Age:      age,
		CheckIn:  CheckIn,
		CheckOut: CheckOut,
	}
}
