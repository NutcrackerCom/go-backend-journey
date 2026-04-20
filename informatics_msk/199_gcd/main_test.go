package main

import "testing"

func TestBasic(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		reference int
		err       error
	}{
		{
			name:      "valid a=9 b=12",
			a:         9,
			b:         12,
			reference: 3,
			err:       nil,
		},
		{
			name:      "valid a=2 b=5",
			a:         2,
			b:         5,
			reference: 1,
			err:       nil,
		},
		{
			name:      "valid a=1 b=9",
			a:         1,
			b:         9,
			reference: 1,
			err:       nil,
		},
		{
			name:      "valid a=27 b=9",
			a:         27,
			b:         9,
			reference: 9,
			err:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				r := Gcd(tt.a, tt.b)
				if r != tt.reference {
					t.Fatalf("got %d, want %d", r, tt.reference)
				}
			})
	}
}
