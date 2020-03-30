package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sc7639/network-cli-tool/internal/app/network-cli/command"

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
			Name:  "ns",
			Usage: "Looks up the Name servers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				nss, err := command.NS(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Print(nss)
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP address for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ips, err := command.IP(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Print(ips)
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Loos up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := command.CNAME(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Print(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mxs, err := command.MX(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Print(mxs)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
