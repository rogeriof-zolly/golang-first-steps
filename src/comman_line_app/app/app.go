package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generates new CLI application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search for IPs and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "ns",
			Usage:  "Search for NSs",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
