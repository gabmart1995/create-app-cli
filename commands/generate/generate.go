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
		Usage:       "generate a directories or files html, CSS and JS empty",
		Description: "generate a directories or files html, CSS and JS empty",
		Subcommands: []*cli.Command{
			{
				Name:        "file",
				Action:      generateFile,
				Aliases:     []string{"f"},
				Description: "generate a file",
				Usage:       "path to generate the file in the specified path",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Required:    true,
						Name:        "path",
						Aliases:     []string{"p"},
						Destination: &pathDestination,
						Usage:       "path to generate the file in the specified path",
					},
				},
			},
			{
				Name:        "directory",
				Aliases:     []string{"d"},
				Action:      generateDirectory,
				Description: "generate a directory in the specified path",
				Usage:       "path to generate the directory in the specified path",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Required:    true,
						Name:        "path",
						Aliases:     []string{"p"},
						Destination: &pathDestination,
						Usage:       "path to generate the directory",
					},
				},
			},
		},
	}
}

/* genera el archivo especificado */
func generateFile(c *cli.Context) error {

	_, err := os.Stat(pathDestination)

	if !os.IsNotExist(err) {
		fmt.Println(models.ColorRed + "el archivo existe en la ubicacion seleccionada")
		return err
	}

	file, err := os.OpenFile(pathDestination, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	fmt.Println(models.ColorGreen + "Archivo creado con éxito")

	return nil
}

/* genera el archivo espcificado */
func generateDirectory(c *cli.Context) error {

	// create the directories
	if err := os.Mkdir(pathDestination, 0755); err != nil {
		return err
	}

	fmt.Println(models.ColorGreen + "Directorio creado con éxito")

	return nil
}
