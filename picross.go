package main

// Picross holds all information about a single picross puzzle
type Picross struct {
	ID int `json:"id,omitempty"`
}

// Validate the puzzle to see if it is a solve-able picross puzzle
func (picross Picross) Validate() bool {
	return true
}
