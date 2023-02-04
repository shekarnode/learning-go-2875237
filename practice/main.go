package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	states := make(map[string]string)
	states["MH"] = "Maharastra"
	states["KL"] = "Kerala"
	states["TN"] = "Tamil Nadu"

	fmt.Println("states :", states)

	fmt.Println("--Deleting a state ---")
	delete(states, "TN")
	fmt.Println("states :", states)
	states["JK"] = "Jammu & Kashmir"

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	stateKeys := make([]string, len(states))
	i := 0
	for k := range states {
		stateKeys[i] = k
		i++
	}
	fmt.Println(stateKeys)
	sort.Strings(stateKeys)
	fmt.Println(stateKeys)

	for i := range stateKeys {
		fmt.Println(states[stateKeys[i]])
	}
}
