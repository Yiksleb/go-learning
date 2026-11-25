package main

import "fmt"

type Player struct {
	Name   string
	Health int
	Level  int
}
type Enemy struct {
	Name   string
	Damage int
}
type Attacker interface {
	Attack()
}

func main() {
	health := 10
	healthptr := &health
	fmt.Println(health)
	fmt.Println(&health)
	fmt.Println(healthptr)
	fmt.Println(*healthptr)
	*healthptr = 75
	fmt.Println(health)
	palyerHealth := 100
	fmt.Println(palyerHealth)
	increaseHealth(&palyerHealth)
	fmt.Println(palyerHealth)
	attack := 50
	defence := 20
	fmt.Println("Attack:", attack, "Defence:", defence)
	improveStats(&attack, &defence)
	fmt.Println("Attack:", attack, "Defence:", defence)
	alex := Player{
		Name:   "Alex",
		Health: 100,
		Level:  25,
	}
	ork := Enemy{
		Name:   "Ork",
		Damage: 25,
	}
	fmt.Println(alex)
	takeDamege(&alex, 25)
	fmt.Println(alex)
	healPlayer(&alex, 10)
	fmt.Println(alex)
	performAttack(ork)
	performAttack(alex)
}
func increaseHealth(health *int) {
	*health += 20

}
func improveStats(attack *int, defence *int) {
	*attack += 10
	*defence += 5
}
func takeDamege(player *Player, damage int) {
	player.Health -= damage
}
func healPlayer(player *Player, amount int) {
	player.Health += amount
}
func (e Enemy) Attack() {
	fmt.Println(e.Name, "attakcs", e.Damage)
}
func (p Player) Attack() {
	fmt.Println(p.Name, "attacks")
}
func performAttack(attacker Attacker) {
	attacker.Attack()
}
