package lms

import (
	"reflect"
	"testing"
)

func TestFindFactors(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "6",
			args: args{6},
			want: []int{1, 2, 3, 6},
		},
		{
			name: "24",
			args: args{24},
			want: []int{1, 2, 3, 4, 6, 8, 12, 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFactors(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
