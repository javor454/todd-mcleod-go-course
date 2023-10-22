package snippets

import (
	"encoding/json"
	"fmt"
	"log"
)

type clovek struct {
	First string
}

func Json() {
	c1 := clovek{First: "jarda"}
	c2 := clovek{First: "hana"}
	cs := []clovek{c1, c2}

	res, err := json.Marshal(cs)
	if err != nil {
		log.Fatalf("error:%v\n", err)
	}

	var p []clovek

	err = json.Unmarshal(res, &p)
	if err != nil {
		log.Fatalf("error:%v\n", err)
	}
	fmt.Println(p)
	fmt.Println("json END------------")
}
