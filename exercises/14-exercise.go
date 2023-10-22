package exercises

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Koko struct {
	Name string
}
type user struct {
	name    string
	age     int
	sayings []string
}

func Ex14() {
	s := []byte("{\"Name\":\"kokosanel\"}")
	var k Koko
	err := json.Unmarshal(s, &k)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(k)

	s2 := []byte("[{\"Name\":\"kokosanel\"},{\"Name\":\"frajerski\"}]")
	var xk []Koko
	err = json.Unmarshal(s2, &xk)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println(xk)

	xb, err := json.Marshal(k)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	fmt.Println(string(xb))

	xb2, err := json.Marshal(xk)
	if err != nil {
		fmt.Println("marshal err:", err)
	}
	fmt.Println(string(xb2))

	//xb3 := []byte("{\"Name\":\"kaja\"}")
	xk2 := []Koko{{"jarda"}, {"frantisek"}}
	err = json.NewEncoder(os.Stdout).Encode(xk2)
	if err != nil {
		fmt.Println("encode err:", err)
	}

	xi := []int{1, 233, 2, 2, 55, 232, 11, 1, 45, -12, 2, 333, 42, 2}
	slices.Sort(xi)
	fmt.Println(xi)

	u1 := user{"jarda", 34, []string{"co je kamo", "cus brasko", "zabite to je"}}
	u2 := user{"karel", 13, []string{"testovani je luxus", "testy", "integracni"}}
	u3 := user{"lukas", 2, []string{"brasko", "mas hlad", "chalecky"}}

	xu1 := []user{u1, u2, u3}
	xu2 := []user{u1, u2, u3}
	slices.SortStableFunc(xu1, func(a, b user) int {
		return cmp.Compare(a.age, b.age)
	})
	slices.SortStableFunc(xu2, func(a, b user) int {
		return cmp.Compare(a.name, b.name)
	})
	for _, v := range xu1 {
		slices.Sort(v.sayings)
	}
	fmt.Println(xu1)
	for _, v := range xu2 {
		slices.Sort(v.sayings)
	}
	fmt.Println(xu2)

	fmt.Println("Exercise 14 END------------")
}
