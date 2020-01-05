// Permutation Generator in Golang
package main

import (
	"fmt"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string

	n = func(testStr []rune, p []string) []string{
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func isPermutation(str1, str2 string) bool {
	// if lengths are not equal - false
	if len(str1) != len(str2) {
		return false
	}

	// sort both strings/runes
	var rune1 RuneSlice = []rune(str1)
	var rune2 RuneSlice = []rune(str2)

	sort.Sort(rune1)
	sort.Sort(rune2)

	//fmt.Println(string(rune1[:]))
	//fmt.Println(string(rune2[:]))

	// compare rune1 and rune 2 by indexes
	for i := 0; i < len(rune1); i++ {
		if rune1[i] != rune2[i] {
			return false
		}
	}

	return true
}

func main() {
	d := permutations("ABCD")
	c := permutations("CADB")

	for i, j := 0,0 ; i < len(d) && j < len(c); i,j = i+1, j+1  {
		fmt.Printf("%v ------ %v = %v\n", d[i], c[j], isPermutation(d[i], c[j]))
	}
}
