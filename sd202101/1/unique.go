package main

import "fmt"

func main() {
	followers := []string{"John", "Richard", "John", "Jane", "Jane", "Alan"}
	unique := make([]string, 0, len(followers))
	m := make(map[string]struct{})
	for _, v := range followers {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}
	fmt.Println(unique)
}
