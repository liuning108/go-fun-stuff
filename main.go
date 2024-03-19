package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name   string
	id     string
	vesion string
	Position
}

type SpecialEntity struct {
	Entity
	specialFiled float64
}

// 1
func main() {

	e := &SpecialEntity{
		specialFiled: 88,
		Entity: Entity{
			name:   "my special entity",
			id:     "my spceilal id ",
			vesion: "1.0",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
	}
	e.name = "my entity is special  "
	e.y = 220
	fmt.Printf("%+v\n", e)

}
