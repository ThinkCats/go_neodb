package main

import (
	"log"
	"os"

	"neodb/src/cli"
)

func main() {
	//	app := &cli.App{
	//		Name:  "neodb",
	//		Usage: "a simple demo database",
	//		Flags: []cli.Flag{
	//			&cli.StringFlag{
	//				Name:    "host",
	//				Aliases: []string{"H"},
	//				Usage:   "db server host address",
	//			},
	//			&cli.StringFlag{
	//				Name:    "port",
	//				Aliases: []string{"P"},
	//				Usage:   "db server port",
	//			},
	//			&cli.StringFlag{
	//				Name:    "user",
	//				Aliases: []string{"U"},
	//				Usage:   "db server user name",
	//			},
	//			&cli.StringFlag{
	//				Name:    "password",
	//				Aliases: []string{"pwd"},
	//				Usage:   "db server password",
	//			},
	//		},
	//		Action: func(_ *cli.Context) error {
	//			fmt.Println("Hello NeoDB")
	//			return nil
	//		},
	//	}

	app := cli.CliConfig()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

}
