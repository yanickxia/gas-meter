package vm

import (
	"errors"

	"github.com/dop251/goja"
)

type VM interface {
	Run(script string) ([]byte, error)
}

type vm struct {
	vm         *goja.Runtime
	ctx        map[string]interface{}
	initScript string
}

func (v *vm) Run(script string) ([]byte, error) {
	if len(v.initScript) != 0 {
		script += v.initScript
	}

	_, err := v.vm.RunString(script)
	if err != nil {
		return nil, err
	}

	call, ok := goja.AssertFunction(v.vm.Get("call"))
	if !ok {
		return nil, errors.New("not found function call")
	}

	if v.ctx != nil {
		if err := v.vm.Set("ctx", v.ctx); err != nil {
			return nil, err
		}
	}

	res, err := call(goja.Undefined())
	if err != nil {
		panic(err)
	}
	ans := res.Export().([]interface{})
	ret := make([]byte, 0, len(ans))
	for i := range ans {
		ret = append(ret, byte(ans[i].(int64)))
	}
	return ret, nil
}

// Builder builder pattern code
type Builder struct {
	vm *vm
}

func NewVmBuilder() *Builder {
	vm := &vm{
		vm: goja.New(),
	}
	b := &Builder{vm: vm}
	return b
}

func (b *Builder) Ctx(ctx map[string]interface{}) *Builder {
	b.vm.ctx = ctx
	return b
}

func (b *Builder) InitScript(initScript string) *Builder {
	b.vm.initScript = initScript
	return b
}

func (b *Builder) Build() VM {
	return b.vm
}
