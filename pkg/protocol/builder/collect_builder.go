package builder

import (
	"errors"

	"github.com/yanickxia/gas-meter/pkg/protocol"
	"github.com/yanickxia/gas-meter/pkg/protocol/utils"
)

const collectCommand = 0x46

type CollectBuilder struct {
	Dest   byte
	Source byte
	Body   []byte
}

func NewCollectBuilder() *CollectBuilder {
	return &CollectBuilder{}
}

func (b *CollectBuilder) WithDest(dest byte) *CollectBuilder {
	b.Dest = dest
	return b
}

func (b *CollectBuilder) WithSource(source byte) *CollectBuilder {
	b.Source = source
	return b
}

func (b *CollectBuilder) WithBody(body []byte) *CollectBuilder {
	b.Body = body
	return b
}

func (b *CollectBuilder) Build() ([]byte, error) {
	if b.Body == nil {
		return nil, errors.New("must have body")
	}

	ret := make([]byte, 0)
	ret = append(ret, protocol.MagicHeader)
	ret = append(ret, byte(4+len(b.Body)), b.Dest, b.Source, collectCommand)
	ret = append(ret, b.Body...)
	ret = append(ret, utils.GenChecksum(ret))
	return ret, nil
}
