package command

import "testing"

func Test_systemCommand_Type(t *testing.T) {
	type fields struct {
		command *command
	}
	tests := []struct {
		name   string
		fields fields
		want   SystemType
	}{
		{
			name:   "test init",
			fields: fields{command: &command{command: 0x00}},
			want:   SystemInitialization,
		},
		{
			name:   "test self check",
			fields: fields{command: &command{command: 0x01}},
			want:   SystemInspection,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &systemCommand{
				command: tt.fields.command,
			}
			if got := s.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
