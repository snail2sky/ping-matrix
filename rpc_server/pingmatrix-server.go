package main

import (
	"database/sql"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"pingmatrix/tools"
	"sort"
	"strings"
)

var authors = make([]*cli.Author, 0)

type arg struct {
	rpcAddr, dbAddr, db, username, password string
}

var db *sql.DB

func ServeForever(a arg) {
	//rpc注册 service
	err := rpc.Register(new(PingMatrix))
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbInfo := strings.Split(a.dbAddr, ":")
	dbHost, dbPort := dbInfo[0], dbInfo[1]
	db = tools.NewDBConnector(a.username, a.password, dbHost, dbPort, a.db)
	defer db.Close()
	server, err := net.Listen("tcp", a.rpcAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("RPC server listen on %s\n", a.rpcAddr)

	defer server.Close()
	for {
		conn, err := server.Accept()
		log.Printf("New connection from %s\n", conn.RemoteAddr().String())
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

func main() {
	authors = append(authors, &cli.Author{Name: "snail", Email: "xuhaoming@itiger.com"})
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "listen",
				Aliases: []string{"l"},
				Value:   "0.0.0.0:9999",
				Usage:   "rpc server serve ip:port",
			},
			&cli.StringFlag{
				Name:    "db-host",
				Aliases: []string{"d"},
				Value:   "127.0.0.1:3306",
				Usage:   "fping info saved",
			},
			&cli.StringFlag{
				Name:    "db-name",
				Aliases: []string{"n"},
				Value:   "ping",
				Usage:   "store host info database",
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Value:   "ping",
				Usage:   "connect database user",
			},
			&cli.StringFlag{
				Name:    "password",
				Value:   "ping",
				Aliases: []string{"p"},
				Usage:   "connect database password",
			},
		},
		Action: func(c *cli.Context) error {
			a := arg{
				c.String("listen"),
				c.String("db-host"),
				c.String("db-name"),
				c.String("username"),
				c.String("password"),
			}
			ServeForever(a)
			return nil
		},
		Authors:     authors,
		Description: "<Server> This program will help you probe network connection status",
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
