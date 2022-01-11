package main

import (
	"log"
	"os"

	"neodb/cli"
)

func main() {
	app := cli.CliConfig()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
