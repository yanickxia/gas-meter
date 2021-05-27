package parse

import (
	"errors"

	"github.com/yanickxia/gas-meter/pkg/protocol"
	"github.com/yanickxia/gas-meter/pkg/protocol/command"
	"github.com/yanickxia/gas-meter/pkg/protocol/data"
)

type responseParse struct {
}

func (r *responseParse) Parse(protocol protocol.RawProtocol) (data.Data, error) {
	responseType := command.NewResponseCommand(protocol.Command()).Type()
	switch responseType {
	case command.ResponseParam:
		return r.parseParam(protocol)
	default:
		return nil, errors.New("not impl")
	}
}

func (r *responseParse) parseParam(protocol protocol.RawProtocol) (data.Data, error) {
	body := protocol.Body()
	chipSelect := body[0]
	address := []byte{body[1], body[2]}
	length := body[3]

	responseParamData := &data.ResponseParamData{
		ChipSelect: chipSelect,
		Address:    address,
		Length:     length,
	}

	responseData := make([][]byte, 0)
	for i := 0; i < int(length); i += 2 {
		responseData = append(responseData, []byte{body[i+4], body[i+5]})
	}

	responseParamData.Data = responseData
	return responseParamData, nil
}
