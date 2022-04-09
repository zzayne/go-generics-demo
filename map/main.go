package main

import "fmt"

type mapKey interface {
	string | int
}

func getMapKeys[K mapKey, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}
	var res []K
	for k := range m {
		res = append(res, k)
	}
	return res
}

func getMapFromSlice[K mapKey, V any](arr []V, fn func(int) K) map[K]V {
	if len(arr) == 0 {
		return nil
	}

	res := make(map[K]V)
	for i, v := range arr {
		res[fn(i)] = v
	}
	return res
}

type user struct {
	ID   string
	Name string
}

func main() {
	users := []user{
		{
			ID:   "111",
			Name: "AAA",
		},
		{
			ID:   "222",
			Name: "BBB",
		},
		{
			ID:   "333",
			Name: "CCC",
		},
	}

	m := getMapFromSlice(users, func(i int) string {
		return users[i].ID
	})
	fmt.Println(m)

	var keys []string
	keys = getMapKeys(m)
	fmt.Println(keys)
	return
}
