package goja

import (
	"reflect"
	"testing"

	"github.com/yanickxia/gas-meter/pkg/vm"
)

func TestCtx(t *testing.T) {
	withContext := vm.NewVmBuilder().Ctx(map[string]interface{}{
		"name": "jack",
		"data": []byte{0x01, 0x02},
	}).Build()

	run, err := withContext.Run(`
function call(){
	return [ctx.name.length , ctx.data.length]
}
`)

	if err != nil {
		t.Errorf("test failed %s", err)
		return
	}

	if !reflect.DeepEqual(run, []byte{0x04, 0x02}) {
		t.Errorf("Build() got = %v, want %v", run, []byte{0x04, 0x02})
	}

}
