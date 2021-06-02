package goja

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/yanickxia/gas-meter/pkg/vm"
)

var (
	testData1, _ = hex.DecodeString("7e0867034604a4002088")
)

func TestGola(t *testing.T) {
	type fields struct {
		Dest byte
		Body []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "test",
			fields: fields{

			},
			want: testData1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bytes, err := vm.NewVmBuilder().Build().Run(`
function call(){
	data = [0x7e, 0x08, 0x67, 0x03, 0x46, 0x04, 0xa4, 0x00, 0x20]
    data.push(checksum(data))
	return data
}

function checksum(data){
	let checksum = 0
	data.slice(2).forEach(function(it){checksum = it + checksum})
	checksum = checksum % 256
	return 256 - checksum
}
`)
			if err != nil {
				t.Errorf("got err %s", err)
				return
			}

			if !reflect.DeepEqual(bytes, tt.want) {
				t.Errorf("Build() got = %v, want %v", bytes, tt.want)
			}

		})
	}
}
