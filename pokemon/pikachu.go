package pokemon

import "fmt"

type (
	IPikachu interface {
		Thunder(IPokemon)
		EvolveToRaichu() *Raichu
	}

	Pikachu struct {
		Pokemon
	}
)

func NewPikachu(name string, maxHP int, attack int, totalAttack int) *Pikachu {
	newPokemon := NewPokemon(name, maxHP, attack, totalAttack)
	return &Pikachu{Pokemon: *newPokemon}
}

func (p *Pikachu) Thunder(target IPokemon) {
	fmt.Println(p.GetName(), "use Thunder")
	p.BasicAttack(target)
}

func (p *Pikachu) EvolveToRaichu() *Raichu {
	raichu := NewRaichu(p.GetName(), p.GetMaxHP()+100, p.GetAttack()+10, 2)
	fmt.Println(p.GetName(), "Evolve to Raichu")
	fmt.Println()
	return raichu
}
