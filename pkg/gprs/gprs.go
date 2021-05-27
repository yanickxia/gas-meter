package gprs

import (
	"bytes"
	"errors"
)

type ModelInfo struct {
	ModelId        []byte
	Phone          []byte
	DynIP          []byte
	ConnectionTime []byte
	RefreshTime    []byte
}

func NewModelInfo(data []byte) (*ModelInfo, error) {
	if len(data) != 18 {
		return nil, errors.New("mode info should 18 bytes")
	}

	reader := bytes.NewReader(data)
	modelId := make([]byte, 2)
	_, _ = reader.Read(modelId)

	phone := make([]byte, 12)
	_, _ = reader.Read(phone)

	dynip := make([]byte, 4)
	_, _ = reader.Read(dynip)

	connTime := make([]byte, 8)
	_, _ = reader.Read(connTime)

	refTime := make([]byte, 8)
	_, _ = reader.Read(refTime)

	return &ModelInfo{
		ModelId:        modelId,
		Phone:          phone,
		DynIP:          dynip,
		ConnectionTime: connTime,
		RefreshTime:    refTime,
	}, nil
}
