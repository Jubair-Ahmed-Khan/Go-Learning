package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	//creating map
	m := make(map[string]string)

	//adding values to map
	m["name"] = "golang"
	m["area"] = "backend"

	//accessing values in map

	// if key doesn't exist, it returns zero value of type
	fmt.Println(m["name"])
	fmt.Println(m["area"])

	//length of map
	fmt.Println(len(m))

	//check if key exists in map
	v, ok := m["age"]

	fmt.Println("value:", v)
	if ok {
		fmt.Println("exists")
	} else {
		fmt.Println("doesn't exist")
	}

	//delete from map
	delete(m, "area")
	fmt.Println(m)

	//delete map
	clear(m)
	fmt.Println("empty map => ", m)

	// map without make keyword
	m2 := map[string]string{"name": "javascript", "area": "frontend"}
	fmt.Println(m2)

	//maps package methods
	// for k, v := range m2 {
	// 	fmt.Println(k, v)
	// }

	for k := range maps.Keys(m2) {
		fmt.Println(k)
	}

	fmt.Println("IsEqual:", maps.Equal(m, m2))
	fmt.Println("Clone:", maps.Clone(m2))

}
