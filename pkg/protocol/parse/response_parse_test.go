package parse

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/yanickxia/gas-meter/pkg/protocol"
	"github.com/yanickxia/gas-meter/pkg/protocol/data"
)

func Test_responseParse_parseParam(t *testing.T) {
	t1, _ := hex.DecodeString("7E2803676604A400200033625AFE5A25200D694168054B0B00037413800262FE46000000002414271041")

	type args struct {
		protocol protocol.RawProtocol
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr bool
	}{
		{
			name: "test parse",
			args: args{protocol: protocol.NewRawProtocol(t1)},
			want: &data.ResponseParamData{
				ChipSelect: 0x04,
				Address:    []byte{0xA4, 0x00},
				Length:     0x20,
				Data: [][]byte{
					{
						0x00, 0x33,
					},
					{
						0x62, 0x5A,
					},
					{
						0xFE, 0x5A,
					},
					{
						0x25, 0x20,
					},
					{
						0x0D, 0x69,
					},
					{
						0x41, 0x68,
					},
					{
						0x05, 0x4B,
					},
					{
						0x0B, 0x00,
					},
					{
						0x03, 0x74,
					},
					{
						0x13, 0x80,
					},
					{
						0x02, 0x62,
					},
					{
						0xFE, 0x46,
					},
					{
						0x00, 0x00,
					},
					{
						0x00, 0x00,
					},
					{
						0x24, 0x14,
					},
					{
						0x27, 0x10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &responseParse{}
			got, err := r.parseParam(tt.args.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}
