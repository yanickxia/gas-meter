package command

// Collect SubCommand

type CollectType string

const (
	CollectPassword CollectType = "password"
	CollectAddress  CollectType = "address"
	CollectBaud     CollectType = "baud"
	CollectNone     CollectType = "None"
	CollectDevice   CollectType = "device"
	CollectParam    CollectType = "param"
	CollectReserve  CollectType = "reserve"
)

type CollectCommand interface {
	Type() CollectType
}

type collectCommand struct {
	*command
}

func NewCollectCommand(command *command) CollectCommand {
	return &collectCommand{command: command}
}

func (c *collectCommand) Type() CollectType {
	cmd := c.command.command << 3 >> 3
	switch cmd {
	case 0x00:
		return CollectPassword
	case 0x01:
		return CollectAddress
	case 0x02:
		return CollectBaud
	case 0x04:
		return CollectDevice
	case 0x06:
		return CollectParam
	default:
		return CollectReserve
	}
}
