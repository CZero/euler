package lms

import "testing"

func TestGetTriangularNumber(t *testing.T) {
	type args struct {
		nth int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3",
			args: args{3},
			want: 6,
		},
		{
			name: "6",
			args: args{5},
			want: 15,
		},
		{
			name: "60",
			args: args{60},
			want: 1830,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTriangularNumber(tt.args.nth); got != tt.want {
				t.Errorf("GetTriangularNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
