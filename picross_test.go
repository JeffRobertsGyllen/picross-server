package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data
//var puzzleData Picross

func TestValidate(t *testing.T) {
	tests := []struct {
		name         string
		expectedExit bool
		picross      Picross
	}{
		{"Empty puzzle", false, Picross{}},
		{"No Squares Filled In", false, Picross{Solution: [][]bool{
			{false, false},
			{false, false},
		}}},
		{"Columns With Different Lengths", false, Picross{Solution: [][]bool{
			{false, true, false},
			{false, false},
			{true, false, true},
		}}},
		{"Valid Small Puzzle", true, Picross{Solution: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		}}},
		// invalid small puzzle - square
		{"Invalid Small Puzzle - Square", false, Picross{Solution: [][]bool{
			{false, true},
			{true, false},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedExit, tt.picross.Validate())
		})
	}
}

//Create puzzle, which validates and fills clue arrays
