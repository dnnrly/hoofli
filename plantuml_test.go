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

func Test_shouldRenderLegend(t *testing.T) {
	type args struct {
		initiatorTypesUsed map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Only unspecified type used",
			args: args{map[string]bool{
				"script":   false,
				"renderer": false,
				"other":    false,
				"":         true,
			}},
			want: false,
		},
		{
			name: "Unspecified and renderer used",
			args: args{map[string]bool{
				"script":   false,
				"renderer": true,
				"other":    false,
				"":         true,
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldRenderLegend(tt.args.initiatorTypesUsed); got != tt.want {
				t.Errorf("shouldRenderLegend() = %v, want %v", got, tt.want)
			}
		})
	}
}
