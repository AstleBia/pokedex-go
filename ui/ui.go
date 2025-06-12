package ui

import (
	"fmt"
	"pokedex-go/pokemon"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var pokemon_name string
var Search bool = true

func Search_pokemon(){
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Type a Pokemons name: ").
				Value(&pokemon_name).
				Placeholder("Pikachu").
				CharLimit(50),
			huh.NewConfirm().
				Title("Search").
				Value(&Search),
		),
	)
	err := form.Run()
	pokemon1 := Pokemon.Buscar_pokemon(pokemon_name)
	var tipos []string
	for _, tipo := range pokemon1.Tipos{
		tipo := tipo.Tipo.Nome
		tipos = append(tipos, tipo)
	}
	if err != nil{
		fmt.Println("Ops...", err)
		return
	}
	tipos_str := strings.Join(tipos,", ")

	dados_pokemon_estilo := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008000")).
		Background(lipgloss.Color("#000000")).
		Bold(true).
		Padding(1).
		Width(24)

	titulo_estilo := lipgloss.NewStyle().
		Inherit(dados_pokemon_estilo).
		Align(lipgloss.Center)

	titulo := titulo_estilo.Render(fmt.Sprintf("%s", strings.ToUpper(pokemon1.Nome)))
	dados := dados_pokemon_estilo.Render(fmt.Sprintf("Id: %v\nPeso: %v\nAltura: %v\nTipo: %v",pokemon1.Id, pokemon1.Peso, pokemon1.Altura, tipos_str))
	bloco := lipgloss.JoinVertical(lipgloss.Right, titulo, dados)
	estilo_bloco := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		BorderBackground(lipgloss.Color("#000000"))
	bloco_final := estilo_bloco.Render(bloco)
	fmt.Println(bloco_final)
	pokemon_name=""
}
