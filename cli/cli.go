package main

import (
	"fmt"
	"log"
	"os"

	cap "github.com/mark-adams/cap-go/cap"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pull"
	app.Usage = "pull CAP alerts"
	app.Action = func(c *cli.Context) error {
		fmt.Println("pulling CAP alerts from NWS")
		feed, err := cap.GetNWSAtomFeed()
		if err != nil {
			return err
		}
		fmt.Printf("%v", feed)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
