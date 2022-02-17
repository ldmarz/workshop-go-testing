package go_testing_test

import (
	"github.com/ldmarz/workshop-go-testing/go_testing"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockPokeDoer struct {
	*http.Response
	error
}

func (m mockPokeDoer) Do(request *http.Request) (*http.Response, error) {
	return m.Response, m.error
}

func Test_GetPokemon_ShouldResponseUsingMockClient(t *testing.T) {
	mockDoer := mockPokeDoer{
		Response: &http.Response{
			Body:       io.NopCloser(strings.NewReader("{\"name\": \"charmeleon\",\"type\": \"fire\"}")),
			StatusCode: 200,
		},
		error: nil,
	}
	pokemonService := go_testing.NewPokemonService("https://pokeapi.co", mockDoer)

	pokemon, err := pokemonService.GetPokemon("charmeleon")

	assert.Nil(t, err)
	assert.Equal(t, pokemon.Name, "charmeleon")
	assert.Equal(t, pokemon.Type, "fire")
}

func Test_GetPokemon_ShouldResponseUsingMockServer(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/v2/pokemon/charmeleon", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"name\": \"charmeleon\",\"type\": \"fire\"}"))
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()

	pokemonService := go_testing.NewPokemonService(srv.URL, http.DefaultClient)

	pokemon, err := pokemonService.GetPokemon("charmeleon")

	assert.Nil(t, err)
	assert.Equal(t, pokemon.Name, "charmeleon")
	assert.Equal(t, pokemon.Type, "fire")
}
