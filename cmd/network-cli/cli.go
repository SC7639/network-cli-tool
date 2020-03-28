package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sc7639/network-cli-tool/internal/app/command"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMES, MX records and Name Servers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ns",
			Usage:  "Looks up the Name servers for a Particular Host",
			Flags:  myFlags,
			Action: command.NS,
		},
		{
			Name:   "ip",
			Usage:  "Looks up the IP address for a particular host",
			Flags:  myFlags,
			Action: command.IP,
		},
		{
			Name:  "cname",
			Usage: "Loos up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
