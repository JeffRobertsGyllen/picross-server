package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test data
var puzzleData Picross

func TestValidate(t *testing.T) {
	tests := []struct {
		name         string
		picross      Picross
		expectedExit bool
	}{
		{"Empty puzzle", puzzleData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedExit, tt.picross.Validate())
		})
	}
}
