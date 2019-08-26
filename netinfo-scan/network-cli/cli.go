package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Initalizing the NewApp method of urfave/cli
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = " Lets you Query IPs, CNAME, Records and Name servers"
	app.Version = "1.0.0"

	// Setting up the Name of the host and a default value
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "mybirdyapp.com",
		},
	}

	// App functionalities for each command
	app.Commands = []cli.Command{
		// ns will look up the Name server for the host
		{
			Name:  "ns",
			Usage: "Looks up the Name server for the Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},

		// Cname will look up for the cononical name for the host
		{
			Name:  "cname",
			Usage: "Looks up the Cononical Name for the host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println(cname)
				return nil
			},
		},

		// IP will look up for the ip address of the given host
		{
			Name:  "ip",
			Usage: "Looks up the IP address for the host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println("Error:", err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},

		// Mail Exchange Records for the given host
		{
			Name:  "mx",
			Usage: "Looks up for Mail-Exchange Records for the given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println("Error:", err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},

		// To look up the ports within the given Host
		{
			Name:  "port",
			Usage: "Looks up for the Ports within the given Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				port, err := net.LookupPort("tcp", "https")
				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println(port)
				return nil
			},
		},
	}

	// Returns if any ERROR from the Arguments
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
