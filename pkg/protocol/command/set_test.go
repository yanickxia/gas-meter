package command

import "testing"

func Test_setCommand_Type(t *testing.T) {
	type fields struct {
		command *command
	}
	tests := []struct {
		name   string
		fields fields
		want   SetType
	}{
		{
			name:   "set password",
			fields: fields{command: &command{0x00}},
			want:   SetPassword,
		},
		{
			name:   "set baud",
			fields: fields{command: &command{0x02}},
			want:   SetBaud,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &setCommand{
				command: tt.fields.command,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
