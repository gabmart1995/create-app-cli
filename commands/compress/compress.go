package compress

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	CompressCommand cli.Command
)

func init() {
	CompressCommand = cli.Command{
		Name:        "compress",
		Aliases:     []string{"c"},
		Usage:       "compress the proyect in zip format",
		Description: "compress the proyect in zip format",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Required:    true,
				DefaultText: "file.zip",
			},
			&cli.PathFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Required: false,
			},
			&cli.PathFlag{
				Name:     "directory",
				Aliases:  []string{"d"},
				Required: false,
				// DefaultText: ".",
			},
		},
		Action: compress,
	}
}

func compress(c *cli.Context) error {

	pwd, _ := os.Getwd()

	nameZip := c.String("name")
	nameFile := c.Path("file")
	pathDirectory := c.Path("directory")
	// recursive := c.Bool("recursive") // directorio recursivo

	// en caso que no venga ningun archivo o directorio
	if len(nameFile) == 0 && len(pathDirectory) == 0 {
		return errors.New("debe especificar un directorio o un archivo a comprimir")
	}

	fmt.Println("Creating zip ...")

	fileZip, err := os.Create(nameZip)

	if err != nil {
		return err
	}

	defer fileZip.Close()

	// creamos el contenedor del zip writer
	zipWriter := zip.NewWriter(fileZip)

	// verificamos si solicita es un directorio
	if len(pathDirectory) > 0 {

		// La funcion manejadora de cada archivo
		readItem := func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if err != nil {
				return err
			}

			// relative path para un solo directorio
			// relativePath := strings.TrimPrefix(path, filepath.Dir(path))
			// fmt.Println(path, filepath.Dir(path), relativePath)

			// generamos el buffer zip
			zipFile, err := zipWriter.Create(path)

			if err != nil {
				return err
			}

			// fmt.Println("linea: 109", path)

			// leemos el buffer archivo
			file, err := os.Open(path)

			if err != nil {
				return err
			}

			defer file.Close()

			// copiamos la info del archivo al zip
			if _, err := io.Copy(zipFile, file); err != nil {
				return err
			}

			return nil
		}

		fmt.Println("Leyendo directorios ...")

		err := filepath.Walk(pathDirectory, readItem)

		if err != nil {
			return err
		}

	} else if len(nameFile) > 0 { // si es un archivo

		fmt.Println("leyendo archivos ...")
		file, err := os.Open(filepath.Join(pwd, nameFile))

		if err != nil {
			return err
		}

		defer file.Close()

		fmt.Println("comprimiendo archivos ...")
		zipFile, err := zipWriter.Create(nameFile)

		if err != nil {
			return err
		}

		// copiamos el buffer del archivo seleccionado
		// hacia el buffer del zip
		if _, err := io.Copy(zipFile, file); err != nil {
			return err
		}

		fmt.Println("archivos comprimidos ...")
	}

	zipWriter.Close() // debemos cerrar el archivo zip con el buffer

	fmt.Println("compresión realizada con éxito")

	return nil
}
