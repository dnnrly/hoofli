package hoofli

import "testing"

func TestTypeToColor(t *testing.T) {
	type args struct {
		strType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Known type",
			args: args{strType: "renderer"},
			want: "blue",
		},
		{
			name: "Unknown type",
			args: args{strType: "spaceship"},
			want: "black",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitiatorTypeToColor(tt.args.strType); got != tt.want {
				t.Errorf("InitiatorTypeToColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
