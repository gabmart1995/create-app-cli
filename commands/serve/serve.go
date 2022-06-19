package serve

import (
	"create-app-cli/helpers"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

/** modulo de creacion de servidor web de desarrollo */

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
				Name:     "directory",
				Aliases:  []string{"t"},
				Required: true,
			},
		},
	}
}

func serve(c *cli.Context) error {

	// recogemos las variables
	address := c.String("address")
	port := c.Int("port")
	directory := c.Path("directory")

	if helpers.FileNotExists(directory) {
		return errors.New("no existe directorio donde mostrar los archivos en la web")
	}

	// creacion del servidor de archivos estaticos
	fileServer := http.FileServer(http.Dir(directory))
	http.Handle("/", http.StripPrefix("/", fileServer))

	message := fmt.Sprintf(`
		Este es un servidor de desarrollo. No usar en producción
		Servidor estatico operando en el puerto %d
		Pulsa Control + C para cerrar el servicio
	`, port)

	fmt.Println(message)

	// si existe algun error desde la plataforma
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil))

	return nil
}
