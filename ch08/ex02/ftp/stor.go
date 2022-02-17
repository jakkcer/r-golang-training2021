package ftp

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func (c *Conn) stor(args []string) {
	if len(args) != 1 {
		c.respond(status501)
		return
	}

	path := filepath.Join(c.rootDir, c.workDir, args[0])
	file, err := os.Create(path)
	if err != nil {
		log.Print(err)
		c.respond(status550)
	}
	c.respond(status150)

	dataConn, err := c.dataConnect()
	if err != nil {
		log.Print(err)
		c.respond(status425)
	}
	defer dataConn.Close()

	_, err = io.Copy(file, dataConn)
	if err != nil {
		log.Print(err)
		c.respond(status426)
		return
	}
	c.respond(status226)
}
