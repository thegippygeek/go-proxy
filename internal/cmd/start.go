package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"

	"github.com/xUnholy/go-proxy/internal/cntlm"
)

func StartCommand() cli.Command {
	return cli.Command{
		Name:        "start",
		Aliases:     []string{""},
		Usage:       "proxy start",
		Description: "Start CNTLM Proxy",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "port, p",
				Value:       3128,
				Usage:       "set custom CNTLM `PORT`",
				Destination: &port,
			},
			cli.BoolFlag{
				Name:        "all, a",
				Usage:       "set all CNTLM config",
				Destination: &setAll,
			},
		},
		Action: func(_ *cli.Context) {
			p := fmt.Sprintf("Listen\t%v", port)
			cntlm.UpdateFile(cntlmFile, p)
			cmds := execute.Command{Cmd: "cntlm", Args: []string{"-g"}}
			_, err := execute.RunCommand(cmds)
			if err != nil {
				fmt.Println("CNTLM Proxy couldn't be started. Is it already running?")
				log.Fatal(err)
			}
			fmt.Println(makeProxyURL(port))
		},
	}
}
