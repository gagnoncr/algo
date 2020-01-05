package main

import (
	"fmt"
	"strings"
)

func unq(element string) {

	element = strings.Replace(element, " ", "", -1)

	encountered := map[string]bool{}
	elements := []string{}


	for _, l := range element {
		elements = append(elements, string(l))
	}

	for v:= range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}

	for key, _ := range encountered {
		result = append(result, key)
	}

	fmt.Printf("%v\n", result)
}

func main() {
	elements := []string{"snake and i", "doggy", "fishiy", "birds"}

	// Remove string duplicates, ignoring order.
	for i := 0; i < len(elements); i++ {
		unq(elements[i])
	}
}
