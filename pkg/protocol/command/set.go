package command

// Set SubCommand

type SetType string

const (
	SetPassword     SetType = "password"
	SetLogicAddress SetType = "logicAddress"
	SetBaud         SetType = "baud"
	SetNone         SetType = "none"
	SetReserve      SetType = "reserve"
	SetParam        SetType = "param"
)

type SetCommand interface {
	Type() SetType
}

type setCommand struct {
	*command
}

func (s *setCommand) Type() SetType {
	cmd := s.command.command << 3 >> 3
	switch cmd {
	case 0x00:
		return SetPassword
	case 0x01:
		return SetLogicAddress
	case 0x02:
		return SetBaud
	case 0x03, 0x05, 0x07:
		return SetNone
	case 0x06:
		return SetParam
	default:
		return SetReserve
	}
}

func NewSetCommand(command *command) SetCommand {
	return &setCommand{command: command}
}
