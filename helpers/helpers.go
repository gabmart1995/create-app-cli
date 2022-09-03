package helpers

import (
	"fmt"
	"os"
)

/* valida si un directorio o archivo existe */
func FileNotExists(name string) bool {

	_, err := os.Stat(name)

	return os.IsNotExist(err)
}

func CreateDirectory(path string) error {
	// create the directories
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}

	fmt.Println("Directorio creado con éxito")

	return nil
}

func CreateFile(path string, content string) error {
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		fmt.Println("el archivo existe en la ubicacion seleccionada")
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return err
	}

	file.WriteString(content)

	defer file.Close()

	return nil
}
