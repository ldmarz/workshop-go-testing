package go_testing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PokeRequester interface {
	Do(*http.Request) (*http.Response, error)
}

type PokemonClient struct {
	baseUrl string
	Doer    PokeRequester
}

func NewPokemonService(baseUrl string, pokeRequester PokeRequester) PokemonClient {
	return PokemonClient{
		baseUrl: baseUrl,
		Doer:    pokeRequester,
	}
}

func (ps PokemonClient) GetPokemon(name string) (Pokemon, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%v/api/v2/pokemon/%v", ps.baseUrl, name), nil)
	res, err := ps.Doer.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("ayuda contextual %w", err)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(res.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
