package command

// Special SubCommand

type SpecialType string

const (
	SpecialRemoteSent SpecialType = "remoteSent"
	SpecialReserve    SpecialType = "reserve"
)

type SpecialCommand interface {
	Type() SpecialType
}

type specialCommand struct {
	*command
}

func NewSpecialCommand(command *command) SpecialCommand {
	return &specialCommand{command: command}
}

func (c *specialCommand) Type() SpecialType {
	cmd := c.command.command << 3 >> 3
	switch cmd {
	case 0x00:
		return SpecialRemoteSent
	default:
		return SpecialReserve
	}
}
