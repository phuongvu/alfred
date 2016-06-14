package main

import "sort"

type By func(p1, p2 *Variant) bool

type variantSorter struct {
	variants []Variant
	by       func(p1, p2 *Variant) bool // Closure used in the Less method.
}

func (by By) Sort(variants []Variant) {
	ps := &variantSorter{
		variants: variants,
		by:       by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

//Implement sort interface
func (s *variantSorter) Len() int {
	return len(s.variants)
}

func (s *variantSorter) Swap(i, j int) {
	s.variants[i], s.variants[j] = s.variants[j], s.variants[i]
}

func (s *variantSorter) Less(i, j int) bool {
	return s.by(&s.variants[i], &s.variants[j])
}
