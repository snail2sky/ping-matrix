package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	inter "pingmatrix/rpc_interface"
	"pingmatrix/tools"
	"sort"
	"time"
)

var authors = make([]*cli.Author, 0)

type arg struct {
	addr, comment, args string
	second              time.Duration
}

func ServeForever(a arg) {
	ip := tools.GetLocalIPByTCPConnect(a.addr)
	log.Println("Local ip is", ip)
	var hostInfo = &inter.HostInfo{IP: ip, Comment: a.comment}
	var response = &inter.RpcResponse{}
	var fPingInfoArr []inter.FPingInfo
	ticker := time.NewTicker(time.Second * a.second)
	rpcConn := NewRpcConn(a.addr)
	log.Printf("Connect rpc server %s", a.addr)
	err := rpcConn.UploadHostInfo(*hostInfo, response)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Will run forever, interval is %v\n", a.second*time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("-------start single loop-------")
			err := rpcConn.GetHostsInfo("", response)
			if err != nil {
				continue
			}
			fPingInfoArr = fPing(ip, a.args, response.AddrArray...)
			err = rpcConn.UploadFPingInfo(fPingInfoArr, response)
			if err != nil {
				continue
			}
			log.Println("-------stop single loop -------")
		}
	}
}

func main() {
	authors = append(authors, &cli.Author{Name: "snail", Email: "xuhaoming@itiger.com"})
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "rpc-addr",
				Aliases: []string{"r"},
				Value:   "127.0.0.1:9999",
				Usage:   "query hosts info, upload local ip and FPing result to host:port",
			},
			&cli.StringFlag{
				Name:    "comment",
				Aliases: []string{"c"},
				Value:   "",
				Usage:   "current machine comment",
			},
			&cli.StringFlag{
				Name:    "fping-args",
				Aliases: []string{"f"},
				Value:   "-q -p 12000 -c 5",
				Usage:   "Provide options arguments for fping",
			},
			&cli.StringFlag{
				Name:    "second",
				Aliases: []string{"s"},
				Value:   "60",
				Usage:   "upload fping info interval second",
			},
		},
		Action: func(c *cli.Context) error {
			a := arg{
				addr:    c.String("rpc-addr"),
				comment: c.String("comment"),
				args:    c.String("fping-args"),
				second:  time.Duration(c.Float64("second")),
			}
			ServeForever(a)
			return nil
		},
		Authors:     authors,
		Description: "<Client> This program will help you probe network connection status",
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
