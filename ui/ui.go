package ui

import (
	"fmt"
	"pokedex-go/pokemon"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func Search_pokemon(){
	var pokemon_name string
	var search bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Type a Pokemons name: ").
				Value(&pokemon_name).
				Placeholder("Pikachu").
				CharLimit(50),
			huh.NewConfirm().
				Title("Search").
				Value(&search),
		),
	)
	err := form.Run()
	pokemon1 := pokemon.Buscar_pokemon(pokemon_name)
	var tipos []string
	for _, tipo := range pokemon1.Tipos{
		tipo := tipo.Tipo.Nome
		tipos = append(tipos, tipo)
	}
	if err != nil || !search{
		fmt.Println("Ops...", err)
		return
	}
	dados_pokemon := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008000")).
		Background(lipgloss.Color("#000000")).
		Bold(true).
		Padding(1)
	texto := dados_pokemon.Render(fmt.Sprintf("Nome: %v\nId: %v\nPeso: %v\nAltura: %v\n", pokemon1.Nome, pokemon1.Id, pokemon1.Peso, pokemon1.Altura))
	fmt.Println(texto)
}