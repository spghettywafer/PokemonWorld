package main

import (
	hospital_pkg "PokemonWorld/hospital"
	"PokemonWorld/pokemon"
	"fmt"
)

func main() {
	fmt.Println("Pokemon World")

	pokemons := []pokemon.IPokemon{
		pokemon.NewPikachu("pika", 100, 15),
		pokemon.NewPikachu("kachu", 120, 10),
		pokemon.NewBulbasaur("bulba", 200, 5),
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

}
