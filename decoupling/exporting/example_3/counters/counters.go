// Package counters provides allert counter support.
package counters

// alterCounter is an UNexported named type
// that contains an integer counter for alerts
type alterCounter int

// New creates and returns values of the unexported type alterCounter.
// DO NOT DO IT. It's really annoying to use.
func New(value int) alterCounter {
	return alterCounter(value)
}
