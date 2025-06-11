package pokemon
import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Tipo struct{
	Nome string `json:"name"`
	Url string `json:"url"`
}

type Tipos struct{
	Slot int `json:"slot"`
	Tipo Tipo `json:"type"`
}

type Pokemon struct{
	Nome string `json:"name"`
	Id int `json:"id"`
	Peso int `json:"weight"`
	Altura int `json:"height"`
	Tipos []Tipos `json:"types"`
}


func Buscar_pokemon(nome_pokemon string) Pokemon{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", nome_pokemon)
	response,err := http.Get(url)
	if err != nil{
		fmt.Println("Opss...", err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil{
		fmt.Println("Opss...",err)
	}
	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil{
		fmt.Println("Opss...",err)
	}
	return pokemon
}