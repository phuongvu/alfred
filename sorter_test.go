package main

import (
	"reflect"
	"strconv"
	"testing"
)

var (
	data = []Variant{
		Variant{
			Title: "Keyboard 2",
			Price: "50.41",
			Grams: 5,
		},
		Variant{
			Title: "Keyboard 3",
			Price: "0.74",
			Grams: 10,
		},
		Variant{
			Title: "Keyboard 1",
			Price: "5.50",
			Grams: 20,
		},
		Variant{
			Title: "Keyboard 4",
			Price: "5",
			Grams: 17,
		},
	}

	sorted = []Variant{
		Variant{
			Title: "Keyboard 3",
			Price: "0.74",
			Grams: 10,
		},
		Variant{
			Title: "Keyboard 4",
			Price: "5",
			Grams: 17,
		},
		Variant{
			Title: "Keyboard 1",
			Price: "5.50",
			Grams: 20,
		},
		Variant{
			Title: "Keyboard 2",
			Price: "50.41",
			Grams: 5,
		},
	}
)

func TestSortByPrice(t *testing.T) {

	price := func(p1, p2 *Variant) bool {
		price1, _ := strconv.ParseFloat(p1.Price, 64)
		price2, _ := strconv.ParseFloat(p2.Price, 64)
		return price1 < price2
	}

	By(price).Sort(data)
	if !expect(data, sorted) {
		t.Errorf("expected %v", sorted)
		t.Errorf("got %v", data)
	}
}

func expect(data interface{}, sorted interface{}) bool {
	return reflect.DeepEqual(data, sorted)
}
