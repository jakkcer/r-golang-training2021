package ftp

import (
	"fmt"
	"log"
	"net"
)

func (c *Conn) enterPassive() (*net.TCPAddr, error) {
	if c.passive != nil {
		c.passive.Close()
		c.passive = nil
	}
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}
	c.passive = lis
	addr := lis.Addr().(*net.TCPAddr)

	return addr, nil
}

func (c *Conn) pasv() {
	addr, err := c.enterPassive()
	if err != nil {
		log.Print(err)
		c.respond(status451)
		return
	}

	hostPort := fmt.Sprintf("127,0,0,1,%d,%d", addr.Port/256, addr.Port%256)
	dataPort, err := dataPortFromHostPort(hostPort)
	if err != nil {
		log.Print(err)
		c.respond(status501)
		return
	}
	c.dataPort = dataPort

	c.respond(fmt.Sprintf(status227, addr.Port/256, addr.Port%256))
}

func (c *Conn) epsv() {
	addr, err := c.enterPassive()
	if err != nil {
		log.Print(err)
		c.respond(status451)
		return
	}

	hostPort := fmt.Sprintf("127,0,0,1,%d,%d", addr.Port/256, addr.Port%256)
	dataPort, err := dataPortFromHostPort(hostPort)
	if err != nil {
		log.Print(err)
		c.respond(status501)
		return
	}
	c.dataPort = dataPort

	c.respond(fmt.Sprintf(status229, addr.Port))
}
