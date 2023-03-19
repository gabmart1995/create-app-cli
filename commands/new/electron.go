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
	Name            string                 `json:"name"`
	Version         string                 `json:"version"`
	Description     string                 `json:"description"`
	Main            string                 `json:"main"`
	Scripts         map[string]string      `json:"scripts"`
	Author          string                 `json:"author"`
	License         string                 `json:"license"`
	Config          map[string]interface{} `json:"config"`
	DevDependencies map[string]string      `json:"devDependencies"`
	Dependencies    map[string]string      `json:"dependencies"`
}

func createElectron(c *cli.Context) error {

	name := c.String("name")
	var pathFiles = path.Join(pwd, name)

	config := configElectron{
		Name:        name,
		Version:     "1.0.0",
		Description: "My Electron application description",
		Main:        "src/index.js",
		Scripts: map[string]string{
			"electron-app": "electron .",
			"start":        "electron-forge start",
			"package":      "electron-forge package",
			"make":         "electron-forge make",
		},
		Author:  "",
		License: "MIT",
		DevDependencies: map[string]string{
			"@electron-forge/cli":            "6.0.0-beta.44",
			"@electron-forge/maker-deb":      "6.0.0-beta.44",
			"@electron-forge/maker-rpm":      "6.0.0-beta.44",
			"@electron-forge/maker-squirrel": "6.0.0-beta.44",
			"@electron-forge/maker-zip":      "6.0.0-beta.44",
			"electron":                       "10.4.7",
		},
		Dependencies: map[string]string{},
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

	if err := os.Mkdir(path.Join(pathFiles, "src", "frontend"), 0755); err != nil {
		return err
	}

	createPackageJSON(path.Join(pathFiles, "package.json"), &config)
	createFile(path.Join(pathFiles, "src", "index.js"), models.GetIndexElectron())
	createFile(path.Join(pathFiles, "src", "frontend", "index.html"), models.GetElectronHTML())

	fmt.Println("Archivos creados con éxito")

	return nil
}

func createPackageJSON(pathFiles string, config *configElectron) error {

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

func createFile(pathFiles string, model string) error {

	file, err := os.OpenFile(
		pathFiles,
		os.O_CREATE|os.O_WRONLY,
		0755,
	)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(model)

	return nil
}
