/**
* Electron Module
 */
package new

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

type configElectron struct {
	Name           string                 `json:"name"`
	Version        int64                  `json:"version"`
	Description    string                 `json:"description"`
	Main           string                 `json:"main"`
	Scripts        map[string]string      `json:"scripts"`
	Author         string                 `json:"author"`
	License        string                 `json:"license"`
	Config         map[string]interface{} `json:"config"`
	DevDependences map[string]string      `json:"devdependences"`
	Dependences    map[string]string      `json:"dependences"`
}

func createElectron(c *cli.Context) error {

	name := c.String("name")

	config := configElectron{
		Name:        name,
		Version:     1.0,
		Description: "",
		Main:        "src/index.js",
		Scripts: map[string]string{
			"electron-app": "electron .",
			"start":        "electron-forge start",
			"package":      "electron-forge package",
			"make":         "electron-forge make",
		},
		Author:  "",
		License: "",
		DevDependences: map[string]string{
			"@electron-forge/cli":       "6.0.0-beta.44",
			"@electron-forge/maker-deb": "6.0.0-beta.44",
			"@electron-forge/maker-rpm": "6.0.0-beta.44",
			"@electron-forge/squirrel":  "6.0.0-beta.44",
			"@electron-forge/maker-zip": "6.0.0-beta.44",
			"electron":                  "10.4.7",
		},
	}

	/*if err := os.Mkdir(path.Join(pwd, name), 0755); err != nil {
		return err
	}*/

	fmt.Println(config)
	// creamos los archivos

	// package.json

	return nil
}
