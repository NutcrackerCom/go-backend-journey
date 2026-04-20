package main

import "testing"

func TestBasic(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		x, y      int
		reference int
		err       error
	}{
		{
			name:      "valid 2 zeros x=2 y=3",
			slice:     []int{0, 0, 0, 0, 5},
			x:         2,
			y:         3,
			reference: 2,
			err:       nil,
		},
		{
			name:      "valid 3 zeros x=2 y=5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         2,
			y:         5,
			reference: 3,
			err:       nil,
		},
		{
			name:      "valid 4 zeros x=1 y=5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         1,
			y:         5,
			reference: 4,
			err:       nil,
		},
		{
			name:      "not valid x=-1 y=5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         -1,
			y:         5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "not valid x=1 y=1",
			slice:     []int{0, 0, 0, 0, 5},
			x:         1,
			y:         1,
			reference: 1,
			err:       nil,
		},
		{
			name:      "not valid x=5 y=5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         5,
			y:         5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "not valid x=5 y=15",
			slice:     []int{0, 0, 0, 0, 5},
			x:         5,
			y:         15,
			reference: 0,
			err:       nil,
		},
		{
			name:      "not valid x=5 y=-5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         5,
			y:         -5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "not valid x=15 y=5",
			slice:     []int{0, 0, 0, 0, 5},
			x:         15,
			y:         5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "valid x=1 y=1",
			slice:     []int{0},
			x:         1,
			y:         1,
			reference: 1,
			err:       nil,
		},
		{
			name:      "valid empty x=1 y=1",
			slice:     []int{},
			x:         1,
			y:         1,
			reference: 0,
			err:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				r := NumberOfZeros(tt.slice, tt.x, tt.y)
				if r != tt.reference {
					t.Fatalf("got %d, want %d", r, tt.reference)
				}
			})
	}
}
