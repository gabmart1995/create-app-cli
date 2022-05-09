package new

import (
	"create-app-cli/models"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

var (
	/* comando new */
	NewCommand cli.Command

	/** integrate boostrap library */
	library string

	/**name of page or application to create*/
	name string

	/* path to work directory*/
	pwd, _ = os.Getwd()
)

func init() {
	NewCommand = cli.Command{
		Name:        "new",
		Usage:       "create a web scafolder web",
		Description: "create a web or app proyect, add library or frameworks css (optional)",
		Action:      create,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Usage:       "-n nombre_sitio",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &name, // el apuntador donde se almacena la variable del tipo
			},
			&cli.StringFlag{
				Name:        "library",
				Usage:       "-l bootstrap.css | materialize.css",
				Aliases:     []string{"l"},
				DefaultText: "basic",
				Destination: &library, // el apuntador donde se almacena la variable del tipo
			},
		},
	}
}

/* Crea los directorios y los archivos de inicio de la biblioteca */
func create(c *cli.Context) error {

	// create the directories
	if err := os.Mkdir(path.Join(pwd, name), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pwd, name, "components"), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pwd, name, "static"), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pwd, name, "static", "css"), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pwd, name, "static", "js"), 0755); err != nil {
		return err
	}

	// se leen los archivos
	data := readFiles(pwd)

	if !writeFiles(data) {
		return errors.New(models.ColorRed + "hubo un problema al crear los archivos")
	}

	fmt.Printf(models.ColorGreen + "Archivos creados con éxito.\n")

	return nil
}

/* lee los modelos de archivos de forma concurrente */
func readFiles(pwd string) map[string]string {

	var data map[string]string = make(map[string]string)
	bootstrap := library == "bootstrap.css"
	materialize := library == "materialize.css"
	basic := library == ""

	data["index.html"] = models.GetHTMLModel(bootstrap, materialize, basic)
	data["index.css"] = models.GetCSSModel()
	data["index.js"] = models.GetJSModel()

	return data
}

/* escribe los archivos de forma concurrente devuelve falso si hay error */
func writeFiles(data map[string]string) bool {

	// channels
	doneHTML := make(chan bool)
	doneCSS := make(chan bool)
	doneJS := make(chan bool)

	go func(doneHTML chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "index.html"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["index.html"])

		doneHTML <- true

	}(doneHTML)

	go func(doneCSS chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "static", "css", "index.css"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["index.css"])

		doneCSS <- true

	}(doneCSS)

	go func(doneJS chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "static", "js", "index.js"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["index.js"])

		doneJS <- true

	}(doneJS)

	result := <-doneCSS && <-doneHTML && <-doneJS

	// cerramos los canales
	close(doneHTML)
	close(doneCSS)
	close(doneJS)

	return result
}
