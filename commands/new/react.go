package new

import (
	"create-app-cli/models"
	"create-app-cli/snipets"
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func createReactStructure(c *cli.Context) error {
	name := c.String("name")
	latest := c.Bool("latest")
	typescript := c.Bool("typescript")

	var config models.Config

	if latest { // versión >= node16
		config = models.Config{
			Name:        name,
			Version:     "1.0.0",
			Description: "React app with create app cli",
			Main:        "src/index.js",
			Scripts: map[string]string{
				"clean":      "rm dist/bundle.js",
				"build-dev":  "webpack --mode development",
				"build-prod": "webpack --mode production",
				"start":      "webpack serve --mode development --open",
			},
			Author:  "",
			Config:  map[string]interface{}{},
			License: "MIT",
			DevDependencies: map[string]string{
				"@babel/core":         "7.x",
				"@babel/preset-env":   "7.x",
				"@babel/preset-react": "7.x",
				"babel-loader":        "9.x",
				"css-loader":          "6.x",
				"style-loader":        "3.x",
				"file-loader":         "6.x",
				"webpack":             "5.x",
				"webpack-cli":         "5.x",
				"webpack-dev-server":  "4.x",
			},
			Dependencies: map[string]string{
				"react":     "18.x",
				"react-dom": "18.x",
			},
		}

		// añade las depencias faltantes si se trabaja con typescript
		if typescript {
			config.DevDependencies["@types/react"] = "18.x"
			config.DevDependencies["@types/react-dom"] = "18.x"
			config.DevDependencies["ts-loader"] = "9.x"
			config.DevDependencies["typescript"] = "5.x"
		}

	} else { // versiones de node8 a node15
		config = models.Config{
			Name:        name,
			Version:     "1.0.0",
			Description: "React app with create app cli",
			Main:        "src/index.js",
			Scripts: map[string]string{
				"clean":      "rm dist/bundle.js",
				"build-dev":  "webpack --mode development",
				"build-prod": "webpack --mode production",
				"start":      "webpack-dev-server --open",
			},
			Author:  "",
			Config:  map[string]interface{}{},
			License: "MIT",
			DevDependencies: map[string]string{
				"@babel/core":         "7.x",
				"@babel/preset-env":   "7.x",
				"@babel/preset-react": "7.x",
				"babel-loader":        "8.x",
				"css-loader":          "3.x",
				"style-loader":        "1.x",
				"file-loader":         "3.x",
				"webpack":             "4.x",
				"webpack-cli":         "3.x",
				"webpack-dev-server":  "3.4.1",
			},
			Dependencies: map[string]string{
				"react":     "18.x",
				"react-dom": "18.x",
			},
		}

		// añade las depencias faltantes si se trabaja con typescript
		if typescript {
			config.DevDependencies["@types/react"] = "18.x"
			config.DevDependencies["@types/react-dom"] = "18.x"
			config.DevDependencies["ts-loader"] = "8.x"
			config.DevDependencies["typescript"] = "4.x"
		}
	}

	var pathFiles = path.Join(pwd, name)

	if err := os.Mkdir(pathFiles, 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pathFiles, "src"), 0755); err != nil {
		return err
	}

	if err := os.Mkdir(path.Join(pathFiles, "dist"), 0755); err != nil {
		return err
	}

	createPackageJSON(path.Join(pathFiles, "package.json"), &config)

	createFile(path.Join(pathFiles, "dist", "index.html"), snipets.GetIndexModel())

	if typescript {
		createFile(path.Join(pathFiles, "src", "index.tsx"), snipets.GetReactModel())

	} else {
		createFile(path.Join(pathFiles, "src", "index.js"), snipets.GetReactModel())

	}

	createFile(path.Join(pathFiles, "src", "style.css"), "")
	createFile(path.Join(pathFiles, ".babelrc"), snipets.GetBabelConfig())

	// webpack config file
	if latest {
		createFile(path.Join(pathFiles, "webpack.config.js"), snipets.GetWebpackConfig5(typescript))

	} else {
		createFile(path.Join(pathFiles, "webpack.config.js"), snipets.GetWebpackConfig(typescript))

	}

	if typescript {
		createFile(path.Join(pathFiles, "tsconfig.json"), snipets.GetTypescriptConfig())
	}

	fmt.Println("Archivos creados con éxito")

	return nil
}
