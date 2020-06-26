package presentation

import "testing"

func Test_namedReturns(t *testing.T) {
	tests := []struct {
		name  string
		wantA string
		wantB string
	}{
		{
			name:  "test namedReturn - pass",
			wantA: "a",
			wantB: "b",
		},
		{
			name:  "test namedReturn - fail",
			wantA: "b",
			wantB: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotA, gotB := namedReturns()

			if gotA != tt.wantA {
				t.Errorf("namedReturns() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("namedReturns() gotB = %v, want %v", gotB, tt.wantB)
			}

		})
	}
}
