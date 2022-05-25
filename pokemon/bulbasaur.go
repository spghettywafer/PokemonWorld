package pokemon

import (
	"fmt"
)

type (
	IBulbasaur interface {
		VineWhip(IPokemon)
	}

	Bulbasaur struct {
		Pokemon
		Monster

		// IShine
	}
)

type Monster struct {
	feet int
}

func NewBulbasaur(name string, maxHP int, attack int, totalAttack int) *Bulbasaur {
	newPokemon := NewPokemon(name, maxHP, attack, totalAttack)
	// monster := Monster{feet: 1}
	bulba := Bulbasaur{
		Pokemon: *newPokemon,
	}
	return &bulba
}

func (p *Bulbasaur) VineWhip(target IPokemon) {
	fmt.Println(p.GetName(), "use VineWhip")
	p.BasicAttack(target)
}

func (p *Bulbasaur) Shine() {
	fmt.Println(p.GetName(), "Shine Bulbasaur")
}
