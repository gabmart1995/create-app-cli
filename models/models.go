package models

type Config struct {
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
