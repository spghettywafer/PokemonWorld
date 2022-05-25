package main

import (
	hospital_pkg "PokemonWorld/hospital"
	"PokemonWorld/pokemon"
	"fmt"
	"math/rand"
	"time"
)

func SimulateAttack(attacker pokemon.IPokemon, attack int, defender pokemon.IPokemon) {
	switch attacker.(type) {
	case pokemon.IRaichu:
		if attack == 1 {
			attacker.(pokemon.IPikachu).Thunder(defender)
		} else if attack == 2 {
			attacker.(pokemon.IRaichu).MegaThunder(defender)
		}
	case pokemon.IPikachu:
		attacker.(pokemon.IPikachu).Thunder(defender)
	case pokemon.IBulbasaur:
		attacker.(pokemon.IBulbasaur).VineWhip(defender)
	}

}

func GetAttackSkill(pokemon pokemon.IPokemon) int {
	//random 1 - totalAttack
	//rand.Intn(max-min) + min
	fmt.Println(pokemon.GetTotalAttack())
	if pokemon.GetTotalAttack() == 1 {
		return pokemon.GetTotalAttack()
	}

	ret := rand.Intn(pokemon.GetTotalAttack()) + 1
	fmt.Println("random att", ret)
	return ret
}

func Battle(firstPokemon pokemon.IPokemon, secondPokemon pokemon.IPokemon) {
	attacker := 0
	for firstPokemon.GetHP() > 0 && secondPokemon.GetHP() > 0 {
		//battle
		//1 -> 0, 1
		//2 -> 1, 0

		if attacker == 0 {
			//1st pokemon attack
			chooseAttack := GetAttackSkill(firstPokemon)
			SimulateAttack(firstPokemon, chooseAttack, secondPokemon)
		} else {
			//2nd pokemon attack
			chooseAttack := GetAttackSkill(secondPokemon)
			SimulateAttack(secondPokemon, chooseAttack, firstPokemon)
		}
		attacker = attacker ^ 1

		// if attacker == 1 {
		// 	attacker = 0
		// } else {
		// 	attacker = 1
		// }
	}

	//win
	if attacker == 1 {
		//1st pokemon win
		fmt.Println(firstPokemon.GetName(), "Win")
	} else {
		//2nd pokemon win
		fmt.Println(secondPokemon.GetName(), "Win")
	}
}

func main() {
	fmt.Println("Pokemon World")
	rand.Seed(time.Now().UnixNano())

	pokemons := []pokemon.IPokemon{
		pokemon.NewPikachu("pika", 100, 15, 1),
		pokemon.NewPikachu("kachu", 120, 10, 1),
		pokemon.NewBulbasaur("bulba", 200, 5, 1),
	}

	pokemons[0].(*pokemon.Pikachu).Thunder(pokemons[1])
	raichu := pokemons[0].(pokemon.IPikachu).EvolveToRaichu()
	pokemons = append(pokemons, raichu)

	pokemons[3].(*pokemon.Raichu).MegaThunder(pokemons[1])

	pokemons[2].(*pokemon.Bulbasaur).VineWhip(pokemons[0])

	pokemons[2].(*pokemon.Bulbasaur).Shine()

	fmt.Println(pokemons[0].GetName(), "HP:", pokemons[0].GetHP(), "/", pokemons[0].GetMaxHP())
	fmt.Println(pokemons[1].GetName(), "HP:", pokemons[1].GetHP(), "/", pokemons[1].GetMaxHP())
	fmt.Println()

	hospital := hospital_pkg.NewHospital()
	pikaHospital := hospital_pkg.NewPikachuHospital()

	hospital.Heal(pokemons[0])
	pikaHospital.Heal(pokemons[1].(pokemon.IPikachu))

	fmt.Println()
	fmt.Println(pokemons[0].GetName(), "HP:", pokemons[0].GetHP(), "/", pokemons[0].GetMaxHP())
	fmt.Println(pokemons[1].GetName(), "HP:", pokemons[1].GetHP(), "/", pokemons[1].GetMaxHP())

	Battle(pokemons[2], pokemons[3])
}
