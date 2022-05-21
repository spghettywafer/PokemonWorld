package hospital

import (
	"PokemonWorld/pokemon"
	"fmt"
)

type (
	IPikachuHospital interface {
		Heal(pokemon.IPikachu)
	}

	PikachuHospital struct{}
)

func NewPikachuHospital() *PikachuHospital {
	return &PikachuHospital{}
}

func (h *PikachuHospital) Heal(target pokemon.IPikachu) {
	target.(pokemon.IPokemon).SetHP(target.(pokemon.IPokemon).GetMaxHP())
	fmt.Println("Heal", target.(pokemon.IPokemon).GetName())
}
