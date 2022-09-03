package generate

import (
	"create-app-cli/helpers"
	"create-app-cli/models"
	"fmt"
	"path"

	"github.com/urfave/cli/v2"
)

/* modulo creador de archivos */

var (
	/* comando new */
	GenerateCommand cli.Command
)

/* constructor del modulo */
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
						Required: true,
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "path to generate the file in the specified path",
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
						Required: true,
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "path to generate the directory",
					},
				},
			},
			{
				Name:        "component",
				Aliases:     []string{"c"},
				Action:      generateComponent,
				Description: "generate a web component in your proyect",
				Usage:       "path to generate a web component in the specified path",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Required: true,
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "path to generate the directory",
					},
					&cli.StringFlag{
						Required: true,
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "name of component",
					},
				},
			},
		},
	}
}

/* genera el archivo especificado */
func generateFile(c *cli.Context) error {

	filePath := c.Path("path")

	if err := helpers.CreateFile(filePath, ""); err != nil {
		return err
	}

	fmt.Println("Archivo creado con éxito")

	return nil
}

/* genera el archivo espcificado */
func generateDirectory(c *cli.Context) error {

	filePath := c.Path("path")

	// create the directories
	if err := helpers.CreateDirectory(filePath); err != nil {
		return err
	}

	return nil
}

func generateComponent(c *cli.Context) error {

	filePath := c.Path("path")
	name := c.String("name")

	if err := helpers.CreateDirectory(filePath); err != nil {
		return err
	}

	model := models.GetModelComponent(name)

	// recorremos el map y generamos los archivos
	for key, value := range model {
		if err := helpers.CreateFile(path.Join(filePath, key), value); err != nil {
			return err
		}
	}

	return nil
}
