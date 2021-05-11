package protocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

var (
	testData1, _ = hex.DecodeString("7e0867034604a4002088")
	testData2, _ = hex.DecodeString("7E2803676604A400200033625AFE5A25200D694168054B0B00037413800262FE46000000002414271041")
)

func Test_rawProtocol_Body(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "test body",
			fields: fields{data: testData1},
			want: []byte{
				0x04, 0xa4, 0x00, 0x20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Body(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Body() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rawProtocol_Command(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "test command",
			fields: fields{data: testData1},
			want:   0x46,
		},
		{
			name:   "test command",
			fields: fields{data: testData2},
			want:   0x66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Command(); got != tt.want {
				t.Errorf("Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rawProtocol_Dest(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "test dest",
			fields: fields{data: testData1},
			want:   0x67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Dest(); got != tt.want {
				t.Errorf("Dest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rawProtocol_Header(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "test header",
			fields: fields{data: testData1},
			want:   MagicHeader,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Header(); got != tt.want {
				t.Errorf("Header() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rawProtocol_Invalid(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test valid",
			fields: fields{
				data: testData1,
			},
			wantErr: false,
		},
		{
			name: "test valid",
			fields: fields{
				data: testData2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if err := r.Invalid(); (err != nil) != tt.wantErr {
				t.Errorf("Invalid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rawProtocol_Length(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "test len",
			fields: fields{data: testData1},
			want:   0x08,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rawProtocol_Source(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "test source",
			fields: fields{data: testData1},
			want:   0x03,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawProtocol{
				data: tt.fields.data,
			}
			if got := r.Source(); got != tt.want {
				t.Errorf("Source() = %v, want %v", got, tt.want)
			}
		})
	}
}
