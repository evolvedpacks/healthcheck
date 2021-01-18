package pinger

import (
	"errors"
	"net"

	"github.com/evolvedpacks/minepong"
)

var ErrEmptyPong = errors.New("empty pong")

type Pinger struct {
	addr string
	conn net.Conn
}

func New(addr string) (p *Pinger, err error) {
	p = new(Pinger)

	p.addr = addr
	p.conn, err = net.Dial("tcp", addr)

	return
}

func PingOnce(addr string) (p *minepong.Pong, err error) {
	pinger, err := New(addr)
	if err != nil {
		return
	}
	defer pinger.Close()

	p, err = pinger.Ping()

	return
}

func (p *Pinger) Ping() (pong *minepong.Pong, err error) {
	if pong, err = minepong.Ping(p.conn, p.addr); err != nil {
		return
	}

	if pong.Version.Name == "" || pong.Version.Protocol == 0 {
		err = ErrEmptyPong
	}

	return
}

func (p *Pinger) Close() error {
	return p.conn.Close()
}
