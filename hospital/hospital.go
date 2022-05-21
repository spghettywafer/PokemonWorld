package hospital

import (
	"PokemonWorld/pokemon"
	"fmt"
)

type (
	IHospital interface {
		Heal()
	}

	Hospital struct {
	}
)

func NewHospital() *Hospital {
	return &Hospital{}
}

func (h *Hospital) Heal(target pokemon.IPokemon) {
	target.SetHP(target.GetMaxHP())
	fmt.Println("Heal", target.GetName())
}
