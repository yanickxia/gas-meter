package gprs

import (
	"encoding/hex"
	"net"
	"reflect"
	"testing"
)

func TestNewModelInfo(t *testing.T) {
	testData1, _ := hex.DecodeString("413432383133393430303532303033000a00000100")
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ModemInfo
		wantErr bool
	}{
		{
			name: "test new",
			args: args{data: testData1},
			want: &ModemInfo{
				ModemId: "41343238",
				Phone:   "13940052003",
				DynIP:   net.IP{10, 0, 0, 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewModelInfo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewModelInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewModelInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
