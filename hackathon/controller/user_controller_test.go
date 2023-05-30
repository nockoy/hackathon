package controller

import "testing"

func TestRegisterUserCheck(t *testing.T) {
	type args struct {
		name string
		age  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常",
			args: args{
				name: "ichiro",
				age:  25,
			},
			want: true,
		},
		{
			name: "名無し",
			args: args{
				name: "",
				age:  26,
			},
			want: false,
		},
		{
			name: "名前が長すぎる",
			args: args{
				name: "01234567890123456789012345678901234567890123456789ro",
				age:  27,
			},
			want: false,
		},
		{
			name: "若すぎる",
			args: args{
				name: "ichiro",
				age:  19,
			},
			want: false,
		},
		{
			name: "81歳以上",
			args: args{
				name: "ichiro",
				age:  81,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterUserCheck(tt.args.name, tt.args.age); got != tt.want {
				t.Errorf("RegisterUserCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
