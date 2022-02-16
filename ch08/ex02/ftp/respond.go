package ftp

import (
	"fmt"
	"log"
)

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status211 = "211 System status, or system help reply."
	status215 = "215 OSX system type."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	// macでは227にしないと返却できない
	status227 = "227 %q is current working directory."
	// status227 = "227 Entering passive mode. 127,0,0,1,%d,%d"
	status229 = "229 Entering extended passive mode. (|||%v|)"
	status230 = "230 User %s logged in, proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status451 = "451 Requested action aborted. Local error in processing."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Command not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

func (c *Conn) EOL() string {
	switch c.dataType {
	case ascii:
		return "\r\n"
	case binary:
		return "\n"
	default:
		return "\n"
	}
}
