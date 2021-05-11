package command

import "testing"

func Test_command_Type(t *testing.T) {
	type fields struct {
		command byte
	}
	tests := []struct {
		name   string
		fields fields
		want   Type
	}{
		{
			name:   "test system",
			fields: fields{command: 0x04},
			want:   System,
		},
		{
			name:   "test set",
			fields: fields{command: 0x20},
			want:   Set,
		},
		{
			name:   "test collect",
			fields: fields{command: 0x40},
			want:   Collect,
		},
		{
			name:   "test revs",
			fields: fields{command: 0x66},
			want:   Response,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &command{
				command: tt.fields.command,
			}
			if got := c.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
