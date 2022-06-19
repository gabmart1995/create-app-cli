package new

import (
	"create-app-cli/helpers"
	"create-app-cli/models"
	"errors"
	"fmt"
	"log"
	"path"

	//"log"
	"os"
	//"path"

	"github.com/urfave/cli/v2"
)

/* modulo de generación de proyectos */

var (
	/* comando new */
	NewCommand cli.Command

	/* path to work directory*/
	pwd, _ = os.Getwd()
)

func init() {
	NewCommand = cli.Command{
		Name:        "new",
		Usage:       "create a web scafolder web",
		Description: "create a web scafolder web",
		Subcommands: []*cli.Command{
			{
				Name:        "static",
				Aliases:     []string{"s"},
				Description: "create a web or app proyect, add library or frameworks css (optional)",
				Action:      create,
				Usage:       "create static web or app proyect",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "-n name",
						Aliases:  []string{"n"},
						Required: true,
						// Destination: &name, // el apuntador donde se almacena la variable del tipo
					},
					&cli.StringFlag{
						Name:        "library",
						Usage:       "-l bootstrap | materialize | basic",
						Aliases:     []string{"l"},
						DefaultText: "bootstrap",
						Value:       "basic", // valor por defecto
						// Destination: &library, // el apuntador donde se almacena la variable del tipo
					},
				},
			},
			{
				Name:        "wordpress",
				Aliases:     []string{"w"},
				Description: "Create a theme or plugin personalized to Wordpress",
				Usage:       "create a theme or plugin to wordpress",
				Action:      createWordpressTheme,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "-n name",
						Aliases:  []string{"n"},
						Required: true,
						// Destination: &name, // el apuntador donde se almacena la variable del tipo
					},
					&cli.StringFlag{
						Name:     "type",
						Usage:    "-t plugin | theme | widget",
						Aliases:  []string{"t"},
						Required: true,
					},
				},
			},
		},
	}
}

/* Crea los directorios y los archivos de inicio de la biblioteca */
func create(c *cli.Context) error {

	// recogemos los valores Args
	name := c.String("name")
	library := c.String("library")

	// create the directories
	if err := os.Mkdir(path.Join(pwd, name), 0755); err != nil {
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
	data := readFiles(pwd, library)

	if !writeFiles(data, name) {
		return errors.New("hubo un problema al crear los archivos")
	}

	fmt.Printf("Archivos creados con éxito.\n")

	return nil
}

/* lee los modelos de archivos de forma concurrente */
func readFiles(pwd string, library string) map[string]string {

	var data map[string]string = make(map[string]string)
	bootstrap := library == "bootstrap"
	materialize := library == "materialize"
	basic := library == "" || library == "basic"

	data["index.html"] = models.GetHTMLModel(bootstrap, materialize, basic)
	data["index.css"] = models.GetCSSModel()
	data["index.js"] = models.GetJSModel()

	return data
}

/* escribe los archivos de forma concurrente devuelve falso si hay error */
func writeFiles(data map[string]string, name string) bool {

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

/* crea una estructura de proyecto para wordpress */
func createWordpressTheme(context *cli.Context) error {

	// fmt.Println("creating wordpress theme")

	// name of theme or plugin type
	name := context.String("name")
	typeAction := context.String("type")

	if typeAction == "theme" {

		if err := os.Mkdir(path.Join(pwd, name), 0755); err != nil {
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

		if err := os.Mkdir(path.Join(pwd, name, "static", "img"), 0755); err != nil {
			return err
		}

		if err := os.Mkdir(path.Join(pwd, name, "includes"), 0755); err != nil {
			return err
		}

		data := readTemplateWordpress()

		if !writeFilesWordpress(data, name) {
			return errors.New("hubo un problema al crear los archivos")
		}

		fmt.Printf("Theme creado con éxito.\n")

	} else if typeAction == "widget" {

		if helpers.FileNotExists("widgets") {

			if err := os.Mkdir(path.Join(pwd, "widgets"), 0755); err != nil {
				return err
			}
		}

		createWidgetWordpress(name)

		fmt.Printf("Widget creados con éxito.\n")

	} else if typeAction == "plugin" {

		if helpers.FileNotExists("plugins") {

			if err := os.Mkdir(path.Join(pwd, "plugins"), 0755); err != nil {
				return err
			}
		}

		createPluginWordpress(name)

		fmt.Printf("Plugin creado con éxito.\n")
	}

	return nil
}

/* lee los templates de wordpress */
func readTemplateWordpress() map[string]string {
	var data = make(map[string]string)

	data["index.php"] = models.GetModelWordpress()
	data["style.css"] = models.GetModelStyleWordpress()
	data["functions.php"] = models.GetModelFunctionsWordpress()

	return data
}

/* escribe los archivos de thema inicial */
func writeFilesWordpress(data map[string]string, name string) bool {

	// channels
	donePHP := make(chan bool)
	doneCSS := make(chan bool)
	doneFunc := make(chan bool)

	go func(doneHTML chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "index.php"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["index.php"])

		donePHP <- true

	}(donePHP)

	go func(doneFunc chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "functions.php"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["functions.php"])

		doneFunc <- true

	}(doneFunc)

	go func(doneHTML chan bool) {
		file, err := os.OpenFile(path.Join(pwd, name, "style.css"), os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		file.WriteString(data["style.css"])

		doneCSS <- true

	}(doneCSS)

	result := <-doneCSS && <-donePHP && <-doneFunc

	// cerramos los canales
	close(donePHP)
	close(doneCSS)
	close(doneFunc)

	return result
}

/* create widget wordpress */
func createWidgetWordpress(name string) {

	file, err := os.OpenFile(path.Join(pwd, "widgets", (name+".php")), os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Fatalln(err)
	}

	file.WriteString(models.GetModelWidget())

	defer file.Close()
}

/* create plugin wordpress */
func createPluginWordpress(name string) {

	file, err := os.OpenFile(path.Join(pwd, "plugins", (name+".php")), os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		log.Fatalln(err)
	}

	file.WriteString(models.GetModelPlugin())

	defer file.Close()
}
