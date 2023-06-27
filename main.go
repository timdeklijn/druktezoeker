package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/timdeklijn/druktezoeker/internal/bikes"
	"github.com/timdeklijn/druktezoeker/internal/getter"
	_ "github.com/timdeklijn/druktezoeker/internal/log"
)

func main() {
	app := &cli.App{
		Name:  "druktezoeker",
		Usage: "Bevraag de Crowdedness API.",
		Commands: []*cli.Command{
			{
				Name:        "bikes",
				Description: "Zoek totaal aantal fietsplaatsen voor een trein",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "trains",
						Usage: "train numbers",
					},
					&cli.StringFlag{
						Name:     "api_key",
						EnvVars:  []string{"APIM_SUBSCRIPTION_KEY"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "host",
						EnvVars:  []string{"HOST"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					trainNumbers := c.StringSlice("trains")
					config, err := getter.NewConfig(c.String("api_key"), c.String("host"))
					if err != nil {
						return err
					}

					return fietsplaatsAggregation(config, trainNumbers)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

// fietsplaatsAggregation aggregates the fietsplaatsen for a given station and train numbers and
// writes to the database.
func fietsplaatsAggregation(config *getter.Config, trainNumbers []string) error {
	crowdedness, err := getter.Crowdedness(config, trainNumbers)
	if err != nil {
		return err
	}

	fietsPlaatsen, err := bikes.AggregateBikes(*crowdedness)
	if err != nil {
		return err
	}
	if err := bikes.WriteBikesToDB(fietsPlaatsen); err != nil {
		return err
	}

	return nil
}
