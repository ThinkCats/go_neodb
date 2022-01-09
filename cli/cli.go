package cli

import (
	"bufio"
	"fmt"
	"neodb/parser"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func ShellLoop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("NeoDB -> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		//handle input
		inputStr := strings.TrimSuffix(input, "\n")
		fmt.Println("Hello,Input:" + inputStr)

		sqlNode, err := parser.ParseNode(inputStr)
		if err != nil {
			fmt.Println("SQL Error: ", inputStr)
		} else {
			tableName := parser.ParseTable(sqlNode)
			fmt.Println("table name:", tableName)
		}

	}
}

func CliConfig() *cli.App {
	app := &cli.App{
		Name:  "neodb",
		Usage: "a simple demo database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"H"},
				Usage:   "db server host address",
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"P"},
				Usage:   "db server port",
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"U"},
				Usage:   "db server user name",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"pwd"},
				Usage:   "db server password",
			},
		},
		Action: func(_ *cli.Context) error {
			//TODO add auth info
			fmt.Println("Start in debug mode ...")
			fmt.Println("WelCome To NeoDB, Start Shell")
			ShellLoop()
			return nil
		},
	}

	return app
}
