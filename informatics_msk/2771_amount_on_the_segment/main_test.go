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
			name:      "valid 1-5 x=2 y=4",
			slice:     []int{1, 2, 3, 4, 5},
			x:         2,
			y:         4,
			reference: 9,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=3 y=5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         3,
			y:         5,
			reference: 12,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=1 y=5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         1,
			y:         5,
			reference: 15,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=-1 y=5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         -1,
			y:         5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=1 y=1",
			slice:     []int{1, 2, 3, 4, 5},
			x:         1,
			y:         1,
			reference: 1,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=5 y=5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         5,
			y:         5,
			reference: 5,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=5 y=15",
			slice:     []int{1, 2, 3, 4, 5},
			x:         5,
			y:         15,
			reference: 0,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=5 y=-5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         5,
			y:         -5,
			reference: 0,
			err:       nil,
		},
		{
			name:      "valid 1-5 x=15 y=5",
			slice:     []int{1, 2, 3, 4, 5},
			x:         15,
			y:         5,
			reference: 0,
			err:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				r := Sum(tt.slice, tt.x, tt.y)
				if r != tt.reference {
					t.Fatalf("got %d, want %d", r, tt.reference)
				}
			})
	}
}
