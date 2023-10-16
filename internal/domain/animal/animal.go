package animal

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	// Se o valor JSON for null, defina Valid como false e retorne
	if string(b) == "null" {
		nf.Valid = false
		return nil
	}

	// Caso contrário, desmarcar o float64 JSON
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	nf.Float64 = f
	nf.Valid = true
	return nil
}

func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	// Se o valor JSON for null, defina Valid como false e retorne
	if string(b) == "null" {
		ns.Valid = false
		return nil
	}

	// Caso contrário, desmarcar a string JSON
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ns.String = s
	ns.Valid = true
	return nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

type Animal struct {
	ID      string
	Name    string
	Breed   string
	Age     int
	Weight  NullFloat64 // Continua usando o tipo personalizado NullFloat64 aqui
	Health  NullString  // Use o tipo personalizado NullString aqui
	Alergic NullString  // Use o tipo personalizado NullString aqui também
}

func NewAnimal(id, name, breed string, age int, weight float64, health, alergic string) *Animal {
	return &Animal{
		ID:      id,
		Name:    name,
		Breed:   breed,
		Age:     age,
		Weight:  NullFloat64{sql.NullFloat64{Float64: weight, Valid: weight != 0}},
		Health:  NullString{sql.NullString{String: health, Valid: health != ""}},
		Alergic: NullString{sql.NullString{String: alergic, Valid: alergic != ""}},
	}
}
