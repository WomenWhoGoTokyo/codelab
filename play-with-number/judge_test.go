package main

import "testing"

func Test_PrimeNumdeterminer(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return '素数です' if you set '5'",
			args: args{
				5,
			},
			want: "素数です",
		},
		{
			name: "return '素数ではありません' if you set '6'",
			args: args{
				6,
			},
			want: "素数ではありません",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PrimeNumdeterminer(tt.args.num); gotResult != tt.want {
				t.Errorf("PrimeNumdeterminer() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
