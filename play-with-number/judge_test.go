package main

import "testing"

func Test_primenumdeterminer(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			// TODO: Add test cases.
			name: "return '素数です' if you set '5'",
			args: args{
				5,
			},
			wantResult: "素数です",
		},
		{
			name: "return '素数ではありません' if you set '6'",
			args: args{
				6,
			},
			wantResult: "素数ではありません",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := primenumdeterminer(tt.args.num); gotResult != tt.wantResult {
				t.Errorf("primenumdeterminer() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
