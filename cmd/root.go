package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/xiaobaiskill/eth_debug/share"
)

func Execute() {
	svc := cli.NewApp()
	svc.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     share.RpcUrl,
			Required: true,
			Usage:    "rpc url(archive node)",
		},
	}
	appendCmdList(svc, call, esatmate)

	err := svc.Run(os.Args)
	if err != nil {
		log.Fatalln("Service Crash")
	}
}

func appendCmdList(app *cli.App, subcmd ...*cli.Command) {
	app.Commands = append(app.Commands, subcmd...)
}
