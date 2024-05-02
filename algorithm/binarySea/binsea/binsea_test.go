package binsea

import (
	"testing"
)

func TestSear2(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"firstTest", args{[]int{1, 2, 3, 5}, 5}, 3},
		// 您可以添加更多测试用例，例如未找到目标的情况
		{"notFoundTest", args{[]int{1, 2, 3, 4}, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sear2(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("Sear1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSear1(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"firstTest", args{[]int{1, 2, 3, 5}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sear1(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("Sear1() = %v, want %v", got, tt.want)
			}
		})
	}
}
