package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {"  hello  world  ", []string{"hello", "world"}},
        {"Charmander Bulbasaur PIKACHU", []string{"charmander", "bulbasaur", "pikachu"}},
        {"   ", []string{}},
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Expected %d words, got %d for input '%s'", len(c.expected), len(actual), c.input)
            continue
        }
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("Expected '%s', got '%s' at index %d", c.expected[i], actual[i], i)
            }
        }
    }
}
