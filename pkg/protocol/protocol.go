package protocol

import "errors"

const (
	MagicHeader = 0x7E
)

type RawProtocol interface {
	Header() byte
	Length() byte
	Dest() byte
	Source() byte
	Command() byte
	Body() []byte
	Checksum() byte
	Invalid() error
}

type rawProtocol struct {
	data []byte
}

func NewRawProtocol(data []byte) RawProtocol {
	return &rawProtocol{data: data}
}

func (r *rawProtocol) Invalid() error {
	dataLen := len(r.data)
	if dataLen < 6 {
		return errors.New("at last 6 bytes")
	}

	if dataLen > 54 {
		return errors.New("data length maximum is 54")
	}

	bodyLen := int(r.Length()) - 4
	if bodyLen > 48 {
		return errors.New("body length maximum is 48")
	}

	if bodyLen+6 != dataLen {
		return errors.New("data length invalid")
	}

	// 将目的地址（含目的地址）与校验和（含校验和）之间的所有数据累加和 MOD 256的结果应等于零
	checksum := 0x0
	for i := 2; i < len(r.data); i++ {
		checksum += int(r.data[i])
	}
	if checksum%256 != 0 {
		return errors.New("checksum failed")
	}

	return nil
}

func (r *rawProtocol) Header() byte {
	return r.data[0]
}

func (r *rawProtocol) Length() byte {
	return r.data[1]
}

func (r *rawProtocol) Dest() byte {
	return r.data[2]
}

func (r *rawProtocol) Source() byte {
	return r.data[3]
}

func (r *rawProtocol) Command() byte {
	return r.data[4]
}

func (r *rawProtocol) Body() []byte {
	return r.data[5 : len(r.data)-1]
}

func (r *rawProtocol) Checksum() byte {
	return r.data[len(r.data)-1]
}
