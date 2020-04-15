package main

import (
	"github.com/urfave/cli"
)

func setupCli() *cli.App {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "CF_BUILD_ID",
				Usage: "The ID of the build in Codefresh",
			},
			&cli.StringFlag{
				Name:  "CF_BUILD_URL",
				Usage: "The URL link to the build in Codefresh",
			},
			&cli.StringFlag{
				Name:  "CF_BUILD_TIMESTAMP",
				Usage: "The timestamp of the build in Codefresh",
			},
			&cli.StringFlag{
				Name:  "CF_BRANCH",
				Usage: "Branch name (or Tag depending on the payload json) of the Git repository of the main pipeline, at the time of execution",
			},
			&cli.StringFlag{
				Name:  "CF_PULL_REQUEST_ID",
				Usage: "The pull request id in Codefresh",
			},
			&cli.StringFlag{
				Name:  "CF_PULL_REQUEST_LABELS",
				Usage: "The labels of pull request (Github and Gitlab only)",
			},
			&cli.StringFlag{
				Name:  "New-Relic-Region",
				Usage: "New Relic Region (US or EU). Leave empty for US",
			},
			&cli.StringFlag{
				Name:  "New-Relic-Insights-Url-Override",
				Usage: "Override the default SAAS New Relic URLs for each region",
			},
			&cli.StringFlag{
				Name:  "Message",
				Usage: "Enter your own custom message to add to the log here",
			},
			&cli.StringFlag{
				Name:     "New-Relic-Account-Id",
				Usage:    "New Relic account id that must be provided",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "X-Insert-Key",
				Usage:    "New Relic Insert API Key that must be provided",
				Required: true,
			},
		},
		Action: sendToNewRelicInsights,
		Name:   "new-relic-sender",
	}

	return app
}
