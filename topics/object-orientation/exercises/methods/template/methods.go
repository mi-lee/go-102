// Declare a new struct type to hold information about a tennis player,
// including the number of matches played and the number won.  Add a method to
// this type that calculates the win ratio for the player.  Create a new
// player, and output the win ratio for them.
//
// Template available at: http://play.golang.org/p/jnBw-jtE3n
package main

// Add your imports here.
import (
	"fmt"
)

// Declare a struct type `player` to maintain information about a player.
type player struct {
	name string
	win  int
	lose int
}

// Declare a method that calculates the win ratio for the player.  Note that
// you'll likely need to convert one or more values to floats, which can be
// done like: float32(intValue)
func (p player) winRatio() float32 {
	return float32(p.win) / float32(p.lose)
}

func main() {
	// Create a new player, and output their win ratio.
	player1 := player{name: "Michelle", win: 5, lose: 2}
	player2 := player{name: "Jess", win: 3, lose: 1}

	// If you're feeling adventurous, try creating a slice of multiple players
	// and iterating over the slice, displaying the player name and win ratio.
	stats := make([]player, 2, 2)
	stats[0] = player1
	stats[1] = player2

	for i := 0; i < 2; i++ {
		fmt.Printf("%s's win ratio is %f \n", stats[i].name, stats[i].winRatio())
	}
}
