package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollDices(t *testing.T) {

	type args struct {
		n int
		d int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"One roll", args{1, 1}, []int{1}},
		{"Six rolls", args{6, 1}, []int{1, 1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RollDices(tt.args.n, tt.args.d); !assert.ElementsMatch(t, tt.want, got) {
				t.Errorf("RollDices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDicesToNumber(t *testing.T) {
	type args struct {
		dices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Three rolls, 3 sides", args{[]int{1, 3, 2}}, 132},
		{"Nine rolls 6 sides", args{[]int{6, 3, 3, 2, 5, 3, 1, 5, 4}}, 633253154},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DicesToNumber(tt.args.dices); got != tt.want {
				t.Errorf("DicesToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
