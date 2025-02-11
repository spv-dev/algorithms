package main

import (
	"fmt"

	bfs "github.com/spv-dev/algorithms/tree/bfs/pkg/bfs_v1"
)

func main() {
	friends := map[string][]string{
		"I":      {"Alice", "Bob", "Clair"},
		"Bob":    {"Anudj", "Peggi"},
		"Alice":  {"Peggi"},
		"Clair":  {"Tomas", "Johnny"},
		"Anudj":  {},
		"Peggi":  {},
		"Tom":    {},
		"Johnny": {"Mark"},
		"Mark":   {"Helen", "Johnny", "Aram"},
		"Helen":  {},
		"Aram":   {},
	}

	fmt.Println(bfs.GetPath(friends, "I", "Aram"))
	fmt.Println(bfs.GetPath(friends, "I", "Tomas"))
	fmt.Println(bfs.GetPath(friends, "I", "Tom1s"))

	city_map := map[string][]string{
		"Home":     {"Park", "School", "Mail"},
		"Park":     {"Home", "Museum", "Cafe"},
		"School":   {"Home", "Library", "Mail"},
		"Mail":     {"Home", "School", "Hospital"},
		"Library":  {"School", "Hospital"},
		"Hospital": {"Library", "Mail", "Office"},
		"Cafe":     {"Park", "Theater", "Office"},
		"Museum":   {"Park", "Shop"},
		"Shop":     {"Museum", "Theater"},
		"Theater":  {"Shop", "Cafe"},
		"Office":   {"Cafe", "Hospital"},
	}
	fmt.Println(bfs.GetPath(city_map, "Home", "Theater"))

}
