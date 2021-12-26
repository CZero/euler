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

func TestFactorial(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Factorial of 5",
			args: args{5},
			want: 120,
		},
		{
			name: "Factorial of 9",
			args: args{9},
			want: 362880,
		},
		{
			name: "Factorial of 10",
			args: args{10},
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.input); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumFactors(t *testing.T) {
	type args struct {
		input  int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "6",
			args: args{6, []int{}},
			want: 4,
		},
		{
			name: "24",
			args: args{24, []int{}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumFactors(tt.args.input, tt.args.primes); got != tt.want {
				t.Errorf("FindNumFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
