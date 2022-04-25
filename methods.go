package main

import "fmt"

type sports struct {
	playername string
	points     int
	event      string
}

func main() {
	player := sports{
		playername: "sridevi",
		points:     30,
		event:      "volleyball",
	}
	player.displayevent()
}
func (s sports) displayevent() {
	fmt.Printf("player %s points%d event:%s", s.playername, s.points, s.event)

}
