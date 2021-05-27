package command

import "testing"

func Test_collectCommand_Type(t *testing.T) {
	type fields struct {
		command *command
	}
	tests := []struct {
		name   string
		fields fields
		want   CollectType
	}{
		{
			name:   "collect pass",
			fields: fields{command: &command{0x00}},
			want:   CollectPassword,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &collectCommand{
				command: tt.fields.command,
			}
			if got := c.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
