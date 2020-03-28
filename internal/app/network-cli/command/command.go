package command

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func NS(c *cli.Context) error {
	ns, err := net.LookupNS(c.String("host"))
	if err != nil {
		return err
	}

	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}

	return nil

}

func IP(c *cli.Context) error {
	ip, err := net.LookupIP(c.String("host"))
	if err != nil {
		return err
	}

	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}

	return nil
}
