package ui
import (
	"fmt"
	"github.com/charmbracelet/huh"
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
	if err != nil || !search{
		fmt.Println("Ops...", err)
		return
	}
	fmt.Printf("Searching for: %s", pokemon_name)
		
}