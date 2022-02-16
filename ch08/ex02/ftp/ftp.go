package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(c *Conn) {
	c.respond(status220)

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}

		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			c.cwd(args)
		// case "EPSV":
		// 	c.epsv()
		case "FEAT":
			c.respond(status211 + " No feature.")
		case "LIST": // ls
			c.list(args)
		// case "PASV":
		// 	c.pasv()
		case "PORT":
			c.port(args)
		case "PWD": // pwd
			c.pwd(args)
		case "QUIT": // close
			c.respond(status221)
			return
		case "RETR": // get
			c.retr(args)
		case "STOR": // put
			c.stor(args)
		case "SYST":
			c.respond(status215)
		case "TYPE":
			c.setDataType(args)
		case "USER":
			c.user(args)
		default:
			c.respond(status502)
		}
	}
	if err := s.Err(); err != nil {
		log.Print(err)
	}
}
