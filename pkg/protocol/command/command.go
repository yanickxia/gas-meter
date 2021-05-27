package command

// Command Def

type Type string

const (
	System   Type = "system"
	Set      Type = "set"
	Collect  Type = "collect"
	Response Type = "response"
	Special  Type = "special"
	Reserve  Type = "reserve"
)

type Command interface {
	Type() Type
}

type command struct {
	command byte
}

func NewCommand(c byte) Command {
	return &command{c}
}

func (c *command) Type() Type {
	cmd := c.command >> 5
	switch cmd {
	case 0x00:
		return System
	case 0x01:
		return Set
	case 0x02:
		return Collect
	case 0x03:
		return Response
	case 0x04:
		return Special
	default:
		return Reserve
	}
}
