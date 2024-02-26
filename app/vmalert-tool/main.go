package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	_ "go.uber.org/automaxprocs"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert-tool/unittest"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/buildinfo"
)

func main() {
	start := time.Now()
	app := &cli.App{
		Name:      "vmalert-tool",
		Usage:     "VMAlert command-line tool",
		UsageText: "More info in https://docs.victoriametrics.com/vmalert-tool.html",
		Version:   buildinfo.Version,
		Commands: []*cli.Command{
			{
				Name:      "unittest",
				Usage:     "Run unittest for alerting and recording rules.",
				UsageText: "More info in https://docs.victoriametrics.com/vmalert-tool.html#Unit-testing-for-rules",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     "files",
						Usage:    "files to run unittest with. Supports an array of values separated by comma or specified via multiple flags.",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "disableAlertgroupLabel",
						Usage:    "disable adding group's Name as label to generated alerts and time series.",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					if failed := unittest.UnitTest(c.StringSlice("files"), c.Bool("disableAlertgroupLabel")); failed {
						return fmt.Errorf("unittest failed")
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Total time: %v", time.Since(start))
}
