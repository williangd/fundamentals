package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Cria um slice com as chaves do mapa
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("keys sem sort:", keys)
	// Ordena as chaves
	sort.Strings(keys)
	fmt.Println("keys com sort:", keys)

	// Printa os estados ordenados
	fmt.Println("\nSorted")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
