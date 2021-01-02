package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var a map[string][]string = make(map[string][]string, 10)
	a["1"] = append(a["1"], "1")
	fmt.Println(a)
}

func Test2(t *testing.T) {
	a := make([]int, 0, 0)
	b := []int{1, 2, 3}
	for {
		a = append(a, b...)
		a = a[len(a)-1:]
	}
}
