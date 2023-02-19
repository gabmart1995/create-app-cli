/**
* Electron Module
 */
package new

import (
	"create-app-cli/models"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

type configElectron struct {
	Name           string                 `json:"name"`
	Version        string                 `json:"version"`
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
	var pathFiles = path.Join(pwd, name)

	config := configElectron{
		Name:        name,
		Version:     "1.0",
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
		Dependences: map[string]string{},
		Config: map[string]interface{}{
			"forge": map[string]interface{}{
				"packagerConfig": map[string]interface{}{},
				"makers": []map[string]interface{}{
					{
						"name":   "@electron-forge/maker-deb",
						"config": map[string]interface{}{},
					},
					{
						"name":   "@electron-forge/maker-rpm",
						"config": map[string]interface{}{},
					},
					{
						"name": "@electron-forge/maker-squirrel",
						"config": map[string]interface{}{
							"name": name,
						},
					},
					{
						"name": "@electron-forge/maker-zip",
						"config": []string{
							"darwin",
						},
					},
				},
			},
		},
	}

	if err := os.Mkdir(pathFiles, 0755); err != nil {
		return err
	}

	// set Path files to new location
	if err := os.Mkdir(path.Join(pathFiles, "src"), 0755); err != nil {
		return err
	}

	createPackageJSON(path.Join(pathFiles, "package.json"), &config)
	createIndexFile(path.Join(pathFiles, "src", "index.js"))

	return nil
}

func createPackageJSON(pathFiles string, config *configElectron) error {

	fmt.Println(pathFiles)

	file, err := os.OpenFile(
		pathFiles,
		os.O_CREATE|os.O_WRONLY,
		0755,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	jsonBytes, _ := json.Marshal(&config)

	// escribimos el archivo
	file.Write(jsonBytes)

	return nil
}

func createIndexFile(pathFiles string) error {

	file, err := os.OpenFile(
		pathFiles,
		os.O_CREATE|os.O_WRONLY,
		0755,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(models.GetIndexElectron())

	return nil
}
