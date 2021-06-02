package gprs

import (
	"bytes"
	"errors"
	"net"
	"time"
)

const (
	RegisterInfoLen = 21
)

type RegisterInfo = ModemInfo

type ModemInfo struct {
	ModemId []byte
	Phone   string
	DynIP   net.IP
}

func (m *ModemInfo) Dest() byte {
	return m.ModemId[3]
}

func NewModelInfo(data []byte) (*ModemInfo, error) {
	if len(data) != 21 {
		return nil, errors.New("mode info should 21 bytes")
	}

	reader := bytes.NewReader(data)
	modelId := make([]byte, 4)
	_, _ = reader.Read(modelId)

	phone := make([]byte, 12)
	_, _ = reader.Read(phone)

	phoneStr := ""
	for i := 0; i < 11; i++ {
		phoneStr += string(phone[i])
	}

	dynip := make([]byte, 4)
	_, _ = reader.Read(dynip)

	return &ModemInfo{
		ModemId: modelId,
		Phone:   phoneStr,
		DynIP:   dynip,
	}, nil
}

type ModemData struct {
	ModemId  uint32
	RecvTime time.Time
	Data     []byte
}
