package data

import "github.com/yanickxia/gas-meter/pkg/protocol/command"

type ResponseParamData struct {
	ChipSelect byte
	Address    []byte
	Length     byte
	Data       [][]byte
}

func (r *ResponseParamData) Command() string {
	return string(command.ResponseParam)
}
