package command

import "testing"

func Test_specialCommand_Type(t *testing.T) {
	type fields struct {
		command *command
	}
	tests := []struct {
		name   string
		fields fields
		want   SpecialType
	}{
		{
			name:   "test remote",
			fields: fields{command: &command{0x00}},
			want:   SpecialRemoteSent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &specialCommand{
				command: tt.fields.command,
			}
			if got := c.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
