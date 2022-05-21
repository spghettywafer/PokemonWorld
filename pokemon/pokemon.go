package pokemon

import (
	"fmt"
	"math/rand"
)

type (
	IShine interface {
		Shine()
	}
)

type (
	IPokemon interface {
		GetName() string
		GetMaxHP() int
		GetHP() int
		SetHP(int)
		GetAttack() int
		BasicAttack(IPokemon)
	}

	Pokemon struct {
		name   string
		maxHP  int
		hp     int
		attack int
	}
)

func NewPokemon(name string, maxHP int, attack int) *Pokemon {
	return &Pokemon{
		name:   name,
		maxHP:  maxHP,
		hp:     maxHP,
		attack: attack,
	}
}

func (p *Pokemon) GetName() string {
	return p.name
}

func (p *Pokemon) GetMaxHP() int {
	return p.maxHP
}

func (p *Pokemon) GetHP() int {
	return p.hp
}

func (p *Pokemon) SetHP(hp int) {
	p.hp = hp
}

func (p *Pokemon) GetAttack() int {
	return p.attack
}

func Attack(p IPokemon, target IPokemon, damage int) {
	fmt.Println(p.GetName(), "deals", damage, "damage to", target.GetName())

	target.SetHP(target.GetHP() - damage)
	fmt.Println(target.GetName(), "becomes", target.GetHP())
	fmt.Println()
}

func (p *Pokemon) BasicAttack(target IPokemon) {
	damage := p.GetAttack() + rand.Intn(2) + 2
	Attack(p, target, damage)
}

func SpecialAttack(p IPokemon, target IPokemon) {
	damage := p.GetAttack() + rand.Intn(10-5) + 10
	Attack(p, target, damage)
}
