package prime

import "testing"

func TestPrime(t *testing.T) {

	tests := []struct {
		arg  int
		want bool
	}{
		{
			arg:  5,
			want: true,
		},
		{
			arg:  6,
			want: false,
		},
	}
	for _, tt := range tests {
		if got := Prime(tt.arg); got != tt.want {
			t.Errorf("Prime() = %v, want %v", got, tt.want)
		}
	}
}
