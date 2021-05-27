package parse

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/yanickxia/gas-meter/pkg/protocol"
	"github.com/yanickxia/gas-meter/pkg/protocol/command"
	"github.com/yanickxia/gas-meter/pkg/protocol/data"
)

type Parse interface {
	Parse(protocol protocol.RawProtocol) (data.Data, error)
}

type parse struct {
	responseParse Parse
}

func (p *parse) Parse(protocol protocol.RawProtocol) (data.Data, error) {
	commandType := command.NewCommand(protocol.Command()).Type()

	switch commandType {
	case command.Response:
		return p.responseParse.Parse(protocol)
	default:
		log.Warningf("got raw protocol %v, do not parse", protocol)
		return nil, errors.New("not need parse")
	}
}
