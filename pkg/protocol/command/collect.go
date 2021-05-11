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
	//TODO
	return ""
}
