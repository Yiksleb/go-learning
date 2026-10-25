package main

import "fmt"

func main() {

	games := map[string]int{
		"Dota2":   5388,
		"WarDogs": 200,
		"Forza 6": 75,
	}
	games = updateGames(games, "Dota2", 12)
	games = updateGames(games, "Minecraft", 40)
	fmt.Println(games)

}

func updateGames(game map[string]int, name string, addedHours int) map[string]int {

	hours, exists := game[name]
	if exists {
		game[name] = hours + addedHours
	} else {
		game[name] = addedHours

	}
	return game
}
