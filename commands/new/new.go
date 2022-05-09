package new

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/urfave/cli/v2"
)

var (
	/* comando new */
	NewCommand cli.Command

	/** tipo de applicacion */
	applicationType string

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
		Aliases:     []string{"n"},
		Usage:       "create a web scafolder web",
		Description: "create a web or app proyect, add library add frameworks css",
		Action:      create,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Usage:       "-n page | application",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &name, // el apuntador donde se almacena la variable del tipo
			},
			&cli.StringFlag{
				Name:        "type",
				Usage:       "-t page | application",
				Aliases:     []string{"t"},
				Required:    true,
				Destination: &applicationType, // el apuntador donde se almacena la variable del tipo
			},
			&cli.StringFlag{
				Name:        "library",
				Usage:       "-l bootstrap.css | materialize.css",
				Aliases:     []string{"l"},
				DefaultText: "bootstrap.css",
				Destination: &library, // el apuntador donde se almacena la variable del tipo
			},
		},
	}
}

/* Crea los directorios y los archivos de inicio de la biblioteca */
func create(c *cli.Context) error {

	var regex []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile(("[{]{2}:link_css[}]{2}")),
		regexp.MustCompile(("[{]{2}:link_script[}]{2}")),
	}

	bootstrap := library == "bootstrap.css" || library == ""
	materialize := library == "materialize.css"

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
	data, err := readFiles(pwd)

	if err != nil {
		return err
	}

	if bootstrap && !materialize {

		var result string

		result = regex[0].ReplaceAllString(
			data["index.html"],
			"<link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\" rel=\"stylesheet\" integrity=\"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC\" crossorigin=\"anonymous\">",
		)

		result = regex[1].ReplaceAllString(
			result,
			"<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM\" crossorigin=\"anonymous\"></script>",
		)

		data["index.html"] = result
	}

	// check the library
	if materialize && !bootstrap {

		var result string

		result = regex[0].ReplaceAllString(
			data["index.html"],
			"<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css\">",
		)

		result = regex[1].ReplaceAllString(
			result,
			"<script src=\"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js\"></script>",
		)

		data["index.html"] = result
	}

	if !writeFiles(data) {
		return errors.New("hubo un problema al crear los archivos")
	}

	fmt.Printf("Archivos creados con éxito.\n")

	return nil
}

/* lee los modelos de archivos de forma concurrente */
func readFiles(pwd string) (map[string]string, error) {

	var data map[string]string = make(map[string]string)

	// channels
	doneHTML := make(chan string)
	doneCSS := make(chan string)
	doneJS := make(chan string)

	// go routines
	go func(doneHTML chan string) {

		file, err := os.OpenFile(path.Join(pwd, "models", "model.html"), os.O_RDONLY, 0755)
		var text string

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		// lee todo el archivo linea a linea y lo acumula en la posicion
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}

		// le pasamos al canal el valor del string
		doneHTML <- text
	}(doneHTML)

	go func(doneCSS chan string) {

		file, err := os.OpenFile(path.Join(pwd, "models", "model.css"), os.O_RDONLY, 0755)
		var text string

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		// lee todo el archivo linea a linea y lo acumula en la posicion
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}

		doneCSS <- text
	}(doneCSS)

	go func(doneJS chan string) {

		file, err := os.OpenFile(path.Join(pwd, "models", "model.js"), os.O_RDONLY, 0755)
		var text string

		if err != nil {
			log.Fatalln(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		// lee todo el archivo linea a linea y lo acumula en la posicion
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}

		doneJS <- text
	}(doneJS)

	// se leen y asigna la informacion de los diversos canales
	data["index.html"] = <-doneHTML
	data["index.css"] = <-doneCSS
	data["index.js"] = <-doneJS

	// cerramos los canales
	close(doneHTML)
	close(doneCSS)
	close(doneJS)

	return data, nil
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
