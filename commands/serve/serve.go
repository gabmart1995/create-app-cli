package serve

import (
	"create-app-cli/helpers"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

/** modulo del servidor de dearrollo */

var (
	ServeCommand cli.Command
)

func init() {
	ServeCommand = cli.Command{
		Name:        "serve",
		Usage:       "run the server in development mode",
		Description: "run the server in development mode",
		Action:      serve, // action
		Aliases:     []string{"S"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Value:       "localhost",
				Aliases:     []string{"A"},
				DefaultText: "localhost",
			},
			&cli.IntFlag{
				Name:        "port",
				Value:       3000,
				Aliases:     []string{"P"},
				DefaultText: "3000",
			},
			&cli.PathFlag{
				Name:        "directory",
				Aliases:     []string{"t"},
				Required:    true,
				DefaultText: ".",
			},
		},
	}
}

/* corre el servidor en modo de desarrollo */
func serve(c *cli.Context) error {

	// recogemos las variables
	address := c.String("address")
	port := c.Int("port")
	directory := c.Path("directory")

	printMessages := func(address string, port int) {

		messageServer := fmt.Sprintf("Servidor estatico operando en la dirección: \"http://%s:%d\"", address, port)
		fmt.Println("Este es un servidor de desarrollo. No usar en producción")
		fmt.Println(messageServer)
		fmt.Println("Pulsa Control + C para cerrar el servicio")
	}

	if helpers.FileNotExists(directory) {
		return errors.New("no existe directorio donde mostrar los archivos en la web")
	}

	// creacion del servidor de archivos estaticos
	fileServer := http.FileServer(http.Dir(directory))
	http.Handle("/", http.StripPrefix("/", fileServer))

	printMessages(address, port)

	// si existe algun error desde la plataforma
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil))

	return nil
}
