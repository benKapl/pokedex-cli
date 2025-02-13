package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/benKapl/pokedex-cli/internal/pokeapi"
)

var (
	cfgTest = &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
)

func TestCommandExplore(t *testing.T) {
	// Must test :
	// - empty location
	// - wrong location
	// - good location : - size of input, inclusion of pokemon

	cases := []struct {
		name        string
		args        []string
		expectedErr error
	}{
		{
			name:        "Empty location",
			args:        []string{},
			expectedErr: errors.New("you must specify a location"),
		},
		{
			name:        "Non existent location",
			args:        []string{"foo"},
			expectedErr: errors.New("this location does not exist"),
		},
		{
			name:        "Valid location",
			args:        []string{"canalave-city-area"},
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Test case: %s", c.name), func(t *testing.T) {
			actualErr := commandExplore(cfgTest, c.args...)

			if actualErr == nil && c.expectedErr == nil {
				return
			}

			if actualErr == nil && c.expectedErr != nil {
				t.Errorf("expected error '%v', but got nil", c.expectedErr)
				return
			}
			if actualErr != nil && c.expectedErr == nil {
				t.Errorf("expected nil, but got error: '%v'", actualErr)
				return
			}
			if actualErr.Error() != c.expectedErr.Error() {
				t.Errorf("errors don't match: '%v' vs '%v'", actualErr, c.expectedErr)
				return
			}
		})

	}

}
