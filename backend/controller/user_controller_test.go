package controller

import "testing"

func TestRegisterUserCheck(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "正常",
			args: args{
				name:  "ichiro",
				email: "ex@gmail.com",
			},
			want: true,
		},
		{
			name: "名無し",
			args: args{
				name: "",
			},
			want: false,
		},
		{
			name: "名前が長すぎる",
			args: args{
				name:  "01234567890123456789012345678901234567890123456789ro",
				email: "ex@gmail.com",
			},
			want: false,
		},
		{
			name: "emailが空欄",
			args: args{
				name:  "ichiro",
				email: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegisterUserCheck(tt.args.name, tt.args.email); got != tt.want {
				t.Errorf("RegisterUserCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
