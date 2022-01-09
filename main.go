package main

import (
	"fmt"
	"log"
	"os"

	"neodb/cli"
)

func main() {

	fmt.Printf("")

	app := cli.CliConfig()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

}
