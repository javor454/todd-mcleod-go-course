package snippets

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type clovecisko struct {
	first string
	age   int
}
type ByAge []clovecisko

func (c clovecisko) String() string {
	return fmt.Sprintf("%s: %d", c.first, c.age)
}

// 1) implementuji metody sort.Interface
func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func Sort() {
	xi := []int{42, 23, 22, 112, 23, 2, 44}
	xs := []string{"asdf", "a", "kk", "ab", "abc"}
	slices.Sort(xi)
	slices.Sort(xs)
	fmt.Println(xi, xs)

	c1 := clovecisko{"jj", 12}
	c2 := clovecisko{"karel", 13}
	c3 := clovecisko{"vasek", 1}
	c4 := clovecisko{"kuba", 442}

	people := []clovecisko{c1, c2, c3, c4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	people2 := []clovecisko{c1, c2, c3, c4}
	slices.SortStableFunc(people2, func(a, b clovecisko) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people2)

	fmt.Println("sort END------------")
}
