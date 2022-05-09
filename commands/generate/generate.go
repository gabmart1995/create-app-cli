package generate

import (
	"create-app-cli/models"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	/* comando new */
	GenerateCommand cli.Command

	/** path */
	pathDestination string
)

func init() {
	GenerateCommand = cli.Command{
		Name:        "generate",
		Aliases:     []string{"g"},
		Usage:       "generate a html, CSS and JS empty",
		Description: "generate a html, CSS and JS empty",
		Action:      generate,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Required:    true,
				Name:        "path",
				Aliases:     []string{"p"},
				Destination: &pathDestination,
				Usage:       "path to generate the file",
			},
		},
	}
}

func generate(c *cli.Context) error {

	_, err := os.Stat(pathDestination)

	if !os.IsNotExist(err) {
		fmt.Println(models.ColorRed + "el archivo existe en la ubicacion seleccionada")
		return err
	}

	writeFile()

	fmt.Printf(models.ColorGreen + "Archivo creado con éxito\n")

	return nil
}

func writeFile() {

	file, err := os.OpenFile(pathDestination, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
}
