package run

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/mrcaelumn/crudwebsocket/api"
	"github.com/mrcaelumn/crudwebsocket/global"
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:  "run",
	Usage: "Run the service",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "socket, s",
			Usage:  "REST API `socket` either as '[tcp://]<address>:<port>' or 'unix://<path>' string",
			EnvVar: "CRUD_WEBSOCKET_SOCKET",
			Value:  "tcp://0.0.0.0:8080",
		},
		cli.StringFlag{
			Name:   "config_file, dff",
			Usage:  "path for file configuration",
			EnvVar: "CONFIG_FILE",
			Value:  "config_file.yaml",
		},
	},
	Action: func(c *cli.Context) error {
		log := log15.New("module", "crudwebsocket")

		var err error
		var listener net.Listener

		// create socket for API server
		socket := c.String("socket")
		if strings.HasPrefix(socket, "unix://") {
			f := strings.TrimPrefix(socket, "unix://")
			if _, err := os.Stat(f); err == nil {
				err = os.Remove(f)
				if err != nil {
					return err
				}
			}
			if listener, err = net.Listen("unix", f); err == nil {
				err = os.Chmod(f, 0770)
			}

		} else {
			if strings.HasPrefix(socket, "tcp://") {
				socket = strings.TrimPrefix(socket, "tcp://")
			}
			listener, err = net.Listen("tcp", socket)
		}
		if err != nil {
			return err
		}

		// capture interrupt signals from OS
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			s := <-sig
			fmt.Println()
			log.Info(fmt.Sprintf("signal %s received", s.String()))
			cancel()
		}()

		if c.String("config_file") != "" {
			os.Setenv("CONFIG_FILE", c.String("config_file"))
			global.CONFIG_FILE = os.Getenv("CONFIG_FILE")
		} else {
			global.CONFIG_FILE = os.Getenv("CONFIG_FILE")
		}

		return api.Run(ctx, listener, log)
	},
}
