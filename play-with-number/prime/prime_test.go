package prime

import "testing"

func TestPrime(t *testing.T) {

	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "is prime",
			args: 5,
			want: true,
		},
		{
			name: "is not prime",
			args: 6,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prime(tt.args); got != tt.want {
				t.Errorf("Prime() = %v, want %v", got, tt.want)
			}
		})
	}
}
