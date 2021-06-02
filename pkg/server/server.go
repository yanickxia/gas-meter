package server

import (
	"bufio"
	"errors"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yanickxia/gas-meter/pkg/gprs"
	"github.com/yanickxia/gas-meter/pkg/protocol"
	"github.com/yanickxia/gas-meter/pkg/protocol/builder"
)

const (
	maxRead = 512
)

type Server interface {
	Run() error
}

type server struct {
	address string
	stop    chan os.Signal
	online  map[string]net.Conn
}

func NewServer(address string) Server {
	return &server{
		address: address,
		stop:    nil,
		online:  map[string]net.Conn{},
	}
}

func (s *server) Run() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Errorf("Error accepting: %s", err.Error())
			continue
		}
		// Handle connections in a new goroutine.
		go func() {
			err := s.handle(conn)
			if err != nil {
				log.Errorf("Error accepting: %s", err.Error())
			}
		}()
	}
}

func (s *server) handle(c net.Conn) error {
	info, err := s.register(c)
	if err != nil {
		return err
	}
	log.Infof("device registed:(id %x, phone %s ip %s)", info.ModemId, info.Phone, info.DynIP)

	conn := &Connection{
		c:    c,
		info: info,
	}

	build, err := builder.NewCollectBuilder().WithDest(info.Dest()).WithSource(0x03).WithBody([]byte{0x04, 0xa4, 0x00, 0x20}).Build()
	rawProtocol := protocol.NewRawProtocol(build)

	for {
		if err := s.command(rawProtocol, conn); err != nil {
			return err
		}

		readProtocol, err := s.readProtocol(conn)
		if err != nil {
			return err
		}
		log.Infof("got data: %v", readProtocol)
	}
}

func (s *server) register(c net.Conn) (*gprs.RegisterInfo, error) {
	shouldRead := gprs.RegisterInfoLen
	bytes := make([]byte, 0)

	for {
		tmp := make([]byte, gprs.RegisterInfoLen)
		read, err := c.Read(tmp)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, tmp[0:read]...)
		log.Debugf("receive bytes %x", tmp[0:read])

		if read <= shouldRead {
			shouldRead -= read
			if shouldRead == 0 {
				break
			}
		} else {
			return nil, errors.New("too much data")
		}
	}

	return gprs.NewModelInfo(bytes)
}

func (s *server) command(protocol protocol.RawProtocol, c *Connection) error {
	//TODO 改成动态载入
	bytes := protocol.Bytes()
	n, err := c.c.Write(bytes)
	if err != nil {
		return err
	}
	log.Debugf("write bytes: %x", bytes[0:n])
	if n != len(bytes) {
		return errors.New("fill data failed")
	}

	return nil
}

func (s *server) readProtocol(c *Connection) (protocol.RawProtocol, error) {
	reader := bufio.NewReader(c.c)
	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if readByte == protocol.MagicHeader {
			break
		}
	}

	l, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	shouldRead := int(l)
	bytes := make([]byte, 0, shouldRead+2)
	bytes = append(bytes, protocol.MagicHeader, l)

	for {
		tmp := make([]byte, maxRead)
		read, err := reader.Read(tmp)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, tmp[0:read]...)
		log.Debugf("receive bytes %x", tmp[0:read])

		if read <= shouldRead {
			shouldRead -= read
			if shouldRead == 0 {
				break
			}
		} else {
			return nil, errors.New("too much data")
		}
	}

	return protocol.NewRawProtocol(bytes), err
}
