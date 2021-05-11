package command

// System SubCommand

type SystemType string

const (
	SystemInitialization   SystemType = "initialization"
	SystemInspection       SystemType = "inspection"
	SystemAdmit            SystemType = "confirmAck"
	SystemDeny             SystemType = "denyAck"
	SystemUnknown          SystemType = "unknownAck"
	SystemPasswordError    SystemType = "passwordError"
	SystemFrameLengthError SystemType = "frameLengthError"
	SystemNone             SystemType = "none"
	SystemReserve          SystemType = "reserve"
)

type SystemCommand interface {
	Type() SystemType
}

type systemCommand struct {
	*command
}

func NewSystemCommand(command *command) SystemCommand {
	return &systemCommand{command: command}
}

func (s *systemCommand) Type() SystemType {
	cmd := s.command.command << 3 >> 3
	switch cmd {
	case 0x00:
		return SystemInitialization
	case 0x01:
		return SystemInspection
	case 0x02:
		return SystemAdmit
	case 0x03:
		return SystemDeny
	case 0x04:
		return SystemUnknown
	case 0x05:
		return SystemPasswordError
	case 0x06:
		return SystemFrameLengthError
	case 0x07:
		return SystemNone
	default:
		return SystemReserve
	}
}
