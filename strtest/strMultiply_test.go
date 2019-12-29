package main

import "testing"

func TestMultiply0043(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"0", args{"0", "0"}, "0"},
		{"1", args{"2", "3"}, "6"},
		{"2", args{"123", "456"}, "56088"},
		{"3", args{"99", "99"}, "9801"},
		{"4", args{"999", "999"}, "998001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("Multiply0043() = %v, want %v", got, tt.want)
			}
		})
	}
}
