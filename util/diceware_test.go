package util

import (
	"reflect"
	"testing"
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
			if got := RollDices(tt.args.n, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RollDices() = %v, want %v", got, tt.want)
			}
		})
	}
}