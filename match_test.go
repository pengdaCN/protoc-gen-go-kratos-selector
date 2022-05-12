package main

import "testing"

func Test_fullMatch(t *testing.T) {
	type args struct {
		selector string
		tags     []string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			args: args{
				selector: "!a|!c",
				tags: []string{
					"c", "afdasf",
				},
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := fullMatch(tt.args.selector, tt.args.tags); gotOk != tt.wantOk {
				t.Errorf("fullMatch() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
