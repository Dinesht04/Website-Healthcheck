package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(ctx context.Context, c *cli.Command) error {
			url := c.Args().Get(0)
			err, resp := GetWebsiteHealth(url)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Response is %v", resp)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
