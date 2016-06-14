package main

import (
	"testing"
	"reflect"
)

var (
	message = &Target{
		[]Product{
			Product{
				"Computer12",
				"Computer",
				[]Variant{
					Variant{
						Title: "Komputer 2",
						Price: "50.41",
						Grams: 7300,
					},
					Variant{
						Title: "Komputer B",
						Price: "5.40",
						Grams: 1500,
					},
				},
			},
			Product{
				"Computer13",
				"Computer",
				[]Variant{
					Variant{
						Title: "Komputer A",
						Price: "17.23",
						Grams: 5700,
					},
				},
			},
			Product{
				"Keyboard2",
				"Keyboard",
				[]Variant{
					Variant{
						Title: "Keyboard LS",
						Price: "11.49",
						Grams: 99,
					},
					Variant{
						Title: "Keyboard AF",
						Price: "31",
						Grams: 500,
					},
				},
			},
			Product{
				"Tower Case",
				"Case",
				[]Variant{
					Variant{
						Title: "Case AS",
						Price: "99.99",
						Grams: 150,
					},
				},
			},
		},
	}

	mapReduceOutput = map[string][]Variant{
		"Computer": []Variant{
			Variant{
				Title: "Computer12 (Komputer B)",
				Price: "5.40",
				Grams: 1500,
			},
			Variant{
				Title: "Computer13 (Komputer A)",
				Price: "17.23",
				Grams: 5700,
			},
			Variant{
				Title: "Computer12 (Komputer 2)",
				Price: "50.41",
				Grams: 7300,
			},
		},
		"Keyboard": []Variant{
			Variant{
				Title: "Keyboard2 (Keyboard LS)",
				Price: "11.49",
				Grams: 99,
			},
			Variant{
				Title: "Keyboard2 (Keyboard AF)",
				Price: "31",
				Grams: 500,
			},
		},
	}

	cartObj = Cart{
		TotalWeight: "7799",
		TotalPrice:  "65.12",
		items: []Variant{
			Variant{
				Title: "Keyboard2 (Keyboard LS)",
				Price: "11.49",
				Grams: 99,
			},
			Variant{
				Title: "Computer12 (Komputer B)",
				Price: "5.40",
				Grams: 1500,
			},
			Variant{
				Title: "Keyboard2 (Keyboard AF)",
				Price: "31",
				Grams: 500,
			},
			Variant{
				Title: "Computer13 (Komputer A)",
				Price: "17.23",
				Grams: 5700,
			},
		},
	}
)

func TestMapReduce(t *testing.T) {
	input := make(chan interface{}, 1)
	input <- message
	close(input)

	results := mapReduce(mapper, reducer, input)

	for c := range results {
		if !reflect.DeepEqual(c, mapReduceOutput) {
			t.Errorf("Expected %v", mapReduceOutput)
			t.Errorf("got %v", c)
		}
	}
}

func TestAddToCart(t *testing.T) {
	result := addToCart(mapReduceOutput)

	if !reflect.DeepEqual(result, cartObj) {
		t.Errorf("Expected %v", cartObj)
		t.Errorf("got %v", result)
	}
}
