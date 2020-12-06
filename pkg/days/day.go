package days

// A Day that can solve the input
type Day interface {
	Part1(Input) (Result, error)
	Part2(Input) (Result, error)
}

// An Input is used to represent the input for the day as a slice of strings
type Input []string

// The Result of the
type Result string
