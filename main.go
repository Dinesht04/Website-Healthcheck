package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dinesht04/health-check/health"
	"github.com/urfave/cli/v3"
)

type Response struct {
	Code    int
	Message string
	Time    time.Duration
}

func CheckHealth(url string) (*Response, error) {
	resp, timeTaken, err := health.GetWebsiteHealth(url)
	if err != nil {
		return nil, err
	}
	response := &Response{}
	response.Code = resp.StatusCode
	response.Time = timeTaken
	if resp.StatusCode == http.StatusOK {
		response.Message = url + " is healthy!"
	} else {
		response.Message = http.StatusText(resp.StatusCode)
	}
	return response, err
}

func main() {
	cmd := &cli.Command{
		Name:      "Website Health Checker",
		Usage:     "Check a Website's health",
		UsageText: "go run main.go check [Your Website's url]",
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "Enter Website's URL",
				Action: func(ctx context.Context, c *cli.Command) error {
					url := c.Args().Get(0)
					resp, err := CheckHealth(url)
					if err != nil {
						return fmt.Errorf("Error - %w", err)
					}

					fmt.Printf("Status Code: %d,\n %s,\n Time Taken: %dms \n", resp.Code, resp.Message, resp.Time.Milliseconds())
					return nil

				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test the tooling",
				Action: func(ctx context.Context, c *cli.Command) error {
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
