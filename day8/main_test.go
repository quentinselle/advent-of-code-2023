package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseMaps(t *testing.T) {
	input := []string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	expected := Maps{
		Directions: []string{"R", "L"},
		Nodes: map[string]Node{
			"AAA": {
				Left:  "BBB",
				Right: "CCC",
			},
			"BBB": {
				Left:  "DDD",
				Right: "EEE",
			},
			"CCC": {
				Left:  "ZZZ",
				Right: "GGG",
			},
			"DDD": {
				Left:  "DDD",
				Right: "DDD",
			},
			"EEE": {
				Left:  "EEE",
				Right: "EEE",
			},
			"GGG": {
				Left:  "GGG",
				Right: "GGG",
			},
			"ZZZ": {
				Left:  "ZZZ",
				Right: "ZZZ",
			},
		},
	}
	actual := parseMaps(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestStepsToReachEnd(t *testing.T) {

	input := []struct {
		Expected int
		Maps     Maps
	}{
		{
			Expected: 2,
			Maps: Maps{
				Directions: []string{"R", "L"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "CCC",
					},
					"BBB": {
						Left:  "DDD",
						Right: "EEE",
					},
					"CCC": {
						Left:  "ZZZ",
						Right: "GGG",
					},
					"DDD": {
						Left:  "DDD",
						Right: "DDD",
					},
					"EEE": {
						Left:  "EEE",
						Right: "EEE",
					},
					"GGG": {
						Left:  "GGG",
						Right: "GGG",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
		{
			Expected: 6,
			Maps: Maps{
				Directions: []string{"L", "L", "R"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "BBB",
					},
					"BBB": {
						Left:  "AAA",
						Right: "ZZZ",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
	}
	for _, test := range input {
		t.Run(fmt.Sprintf("expected_steps_%d", test.Expected), func(t *testing.T) {
			actual := test.Maps.StepsToReachEnd()
			if actual != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, actual)
			}
		})
	}
}

func BenchmarkStepsToReachEnd(b *testing.B) {

	input := []struct {
		Expected int
		Maps     Maps
	}{
		{
			Expected: 2,
			Maps: Maps{
				Directions: []string{"R", "L"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "CCC",
					},
					"BBB": {
						Left:  "DDD",
						Right: "EEE",
					},
					"CCC": {
						Left:  "ZZZ",
						Right: "GGG",
					},
					"DDD": {
						Left:  "DDD",
						Right: "DDD",
					},
					"EEE": {
						Left:  "EEE",
						Right: "EEE",
					},
					"GGG": {
						Left:  "GGG",
						Right: "GGG",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
		{
			Expected: 6,
			Maps: Maps{
				Directions: []string{"L", "L", "R"},
				Nodes: map[string]Node{
					"AAA": {
						Left:  "BBB",
						Right: "BBB",
					},
					"BBB": {
						Left:  "AAA",
						Right: "ZZZ",
					},
					"ZZZ": {
						Left:  "ZZZ",
						Right: "ZZZ",
					},
				},
			},
		},
	}
	for _, test := range input {
		b.Run(fmt.Sprintf("expected_steps_%d", test.Expected), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.Maps.StepsToReachEnd()
			}
		})
	}
}

func TestStepsToReachEndGhost(t *testing.T) {

	input := Maps{
		Directions: []string{"L", "R"},
		Nodes: map[string]Node{
			"11A": {
				Left:  "11B",
				Right: "XXX",
			},
			"11B": {
				Left:  "XXX",
				Right: "11Z",
			},
			"11Z": {
				Left:  "11B",
				Right: "XXX",
			},
			"22A": {
				Left:  "22B",
				Right: "XXX",
			},
			"22B": {
				Left:  "22C",
				Right: "22C",
			},
			"22C": {
				Left:  "22Z",
				Right: "22Z",
			},
			"22Z": {
				Left:  "22B",
				Right: "22B",
			},
			"XXX": {
				Left:  "XXX",
				Right: "XXX",
			},
		},
	}
	expected := 6
	actual := input.StepsToReachEndGhost()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func BenchmarkStepsToReachEndGhost(b *testing.B) {

	input := Maps{
		Directions: []string{"L", "R"},
		Nodes: map[string]Node{
			"11A": {
				Left:  "11B",
				Right: "XXX",
			},
			"11B": {
				Left:  "XXX",
				Right: "11Z",
			},
			"11Z": {
				Left:  "11B",
				Right: "XXX",
			},
			"22A": {
				Left:  "22B",
				Right: "XXX",
			},
			"22B": {
				Left:  "22C",
				Right: "22C",
			},
			"22C": {
				Left:  "22Z",
				Right: "22Z",
			},
			"22Z": {
				Left:  "22B",
				Right: "22B",
			},
			"XXX": {
				Left:  "XXX",
				Right: "XXX",
			},
		},
	}
	for i := 0; i < b.N; i++ {
		input.StepsToReachEndGhost()
	}
}
