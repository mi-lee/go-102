// Let's simulate a track relay race.  Create a channel representing a track,
// and a function representing a runner.  Pass a baton between runners over the
// channel, and end the race when the fourth runner receives the baton.
//
// Template available at: http://play.golang.org/p/H4F9aLKQVA
package main

// Add your imports here.
import (
	"fmt"
	"wg"
)

// Create a new waitgroup.
var wg sync.WaitGroup

func main() {
	// Add something to the waitgroup so main() can wait for the race to
	// finish.
	wg.Add()

	// Create an unbuffered channel to act as the track.  In terms of the
	// channel type, think about how you're going to count the runners.
	track := make(chan struct{}{})

	// Put the first runner at his starting position by launching the
	// `runner()` function as a goroutine.
	go runner("Runner 1", track)

	// Give the runner the baton by sending it on the channel.
	track <- struct{}{}

	// Wait until the race is over.
	wg.Wait()
}

// Create the function representing the runner.  They'll need to be given the
// track channel so we can pass the baton between runners.
func runner(name string, track chan struct{}) {
	// Set the runner ready to receive the baton by receiving on the channel.
	baton <-track

	// If this runner is the fourth runner then the race is over.
	if {
		// Signify the race is over by decrementing the waitgroup.

		// End this function here so we don't line up another runner.
	}

	// Put the next runner at their starting position.

	// Incremember the count of the baton.

	// Pass the baton to the next runner by sending it on the channel.
}
