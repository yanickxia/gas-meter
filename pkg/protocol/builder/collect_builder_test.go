package builder

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var (
	testData1, _ = hex.DecodeString("7e0867034604a4002088")
)

func TestCollectBuilder_Build(t *testing.T) {
	type fields struct {
		Dest   byte
		Source byte
		Body   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Dest:   0x67,
				Source: 0x03,
				Body:   []byte{0x04, 0xa4, 0x00, 0x20},
			},
			want:    testData1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &CollectBuilder{
				Dest:   tt.fields.Dest,
				Source: tt.fields.Source,
				Body:   tt.fields.Body,
			}
			got, err := b.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}
