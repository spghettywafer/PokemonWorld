package pokemon

import "fmt"

type (
	IBulbasaur interface {
		VineWhip(IPokemon)
	}

	Bulbasaur struct {
		Pokemon
	}
)

func NewBulbasaur(name string, maxHP int, attack int) *Bulbasaur {
	newPokemon := NewPokemon(name, maxHP, attack)
	return &Bulbasaur{Pokemon: *newPokemon}
}

func (p *Bulbasaur) VineWhip(target IPokemon) {
	fmt.Println(p.GetName(), "use VineWhip")
	p.BasicAttack(target)
}

func (p *Bulbasaur) Shine() {
	fmt.Println(p.GetName(), "Shine Bulbasaur")
}
