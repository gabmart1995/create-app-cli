package helpers

import "os"

/* valida si un direcotrio o archivo existe */
func FileNotExists(name string) bool {

	_, err := os.Stat(name)

	return os.IsNotExist(err)
}
