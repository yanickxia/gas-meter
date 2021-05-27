package command

// Response SubCommand

type ResponseType string

const (
	ResponsePassword ResponseType = "password"
	ResponseAddress  ResponseType = "address"
	ResponseBaud     ResponseType = "baud"
	ResponseReserve  ResponseType = "reserve"
	ResponseDevice   ResponseType = "device"
	ResponseParam    ResponseType = "param"
)

type ResponseCommand interface {
	Type() ResponseType
}

type responseCommand struct {
	command byte
}

func NewResponseCommand(command byte) ResponseCommand {
	return &responseCommand{command: command}
}

func (c *responseCommand) Type() ResponseType {
	cmd := c.command << 3 >> 3
	switch cmd {
	case 0x00:
		return ResponsePassword
	case 0x01:
		return ResponseAddress
	case 0x02:
		return ResponseBaud
	case 0x04:
		return ResponseDevice
	case 0x06:
		return ResponseParam
	default:
		return ResponseReserve
	}
}
