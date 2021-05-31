package server

import (
	"net"

	"github.com/yanickxia/gas-meter/pkg/gprs"
)

type Connection struct {
	c    net.Conn
	info *gprs.RegisterInfo
}
