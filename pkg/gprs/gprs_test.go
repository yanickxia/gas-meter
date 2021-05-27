package gprs

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestNewModelInfo(t *testing.T) {
	testData1, _ := hex.DecodeString("A42813940052003")
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ModelInfo
		wantErr bool
	}{
		{
			name: "test new",
			args: args{data: testData1},
			want: &ModelInfo{
				ModelId:        nil,
				Phone:          nil,
				DynIP:          nil,
				ConnectionTime: nil,
				RefreshTime:    nil,
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
