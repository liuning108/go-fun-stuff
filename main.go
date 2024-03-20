package main

import (
	"fmt"
)

type Player struct {
	HP int
}

func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage New.HP -> ", p.HP)
}

func main() {
	play := &Player{HP: 100}
	play.TakeDamage(10)
	fmt.Printf("%+v\n", play)
}
