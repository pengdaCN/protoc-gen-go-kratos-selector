package pkg

import "testing"

func Test_combineOperation(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				op: "/api.user.v1.User/Update",
			},
			want: "/api.user.v1.User/Update",
		},
		{
			args: args{
				op: "/api.user.v1.User/update",
			},
			want: "/api.user.v1.User/Update",
		},
		{
			args: args{
				"/api.user.v1.User/u",
			},
			want: "/api.user.v1.User/U",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineOperation(tt.args.op); got != tt.want {
				t.Errorf("CombineOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
