package main

import "testing"

func Test_match(t *testing.T) {
	type args struct {
		selector string
		name     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				selector: "!no-jwt",
				name:     "jwt",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.selector, tt.args.name); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
