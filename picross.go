package main

// Picross holds all information about a single picross puzzle
type Picross struct {
	ID       int      `json:"id,omitempty"`
	Solution [][]bool `json:"solution,omitempty"`
}

// Validate the puzzle to see if it is a solve-able picross puzzle
func (picross Picross) Validate() bool {
	var notEmpty, columnsEqual bool = false, true
	for i := 0; i < len(picross.Solution); i++ {
		columnsEqual = columnsEqual && len(picross.Solution[0]) == len(picross.Solution[i])
		for j := 0; j < len(picross.Solution[i]); j++ {
			notEmpty = notEmpty || picross.Solution[i][j]
		}
	}
	return notEmpty && columnsEqual
}
