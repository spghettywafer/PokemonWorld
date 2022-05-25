package pokemon

import "fmt"

type (
	IRaichu interface {
		MegaThunder(IPokemon)
	}
	Raichu struct {
		Pikachu
	}
)

func NewRaichu(name string, maxHP int, attack int, totalAttack int) *Raichu {
	newPikachu := NewPikachu(name, maxHP, attack, totalAttack)
	return &Raichu{Pikachu: *newPikachu}
}

func (p *Raichu) MegaThunder(target IPokemon) {
	fmt.Println(p.GetName(), "use Mega Thunder")
	SpecialAttack(p, target)
}
