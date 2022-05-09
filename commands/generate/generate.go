package generate

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	/* comando new */
	GenerateCommand cli.Command

	/** tipo de generacion */
	generateType string

	/* path to work directory*/
	pwd, _ = os.Getwd()
)

func init() {
	GenerateCommand = cli.Command{
		Name:        "generate",
		Aliases:     []string{"g"},
		Usage:       "generate a web asset HTML, CSS and JS",
		Description: "generate a web asset HTML, CSS and JS",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required:    true,
				Name:        "type",
				DefaultText: "html",
				Aliases:     []string{"t"},
			},
		},
	}
}
