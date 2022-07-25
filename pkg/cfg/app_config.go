package cfg

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	// EnvProd represents the environment for production
	EnvProd = "prod"
	// EnvDev represents the environment for development
	EnvDev = "dev"
	// EnvStage represents the environment for staging
	EnvStage = "stage"
)

// AppCfg application configuration
type AppCfg struct {
	Server ServerConfig `yaml:"server"`
	Data   DataConfig   `yaml:"data"`
}

// ServerConfig holds configuration information for the application http server
type ServerConfig struct {
	Addr string `yaml:"addr"`
}

// DataConfig holds configuration information for which data files should be loaded
type DataConfig struct {
	Load string `yaml:"load"`
}

// LoadApplicationConfig for the spec. environment
func LoadApplicationConfig(environment string) (app AppCfg, err error) {
	filepath := path.Join("config", fmt.Sprintf("app.%s.yml", environment))

	if err = loadYamlFile(filepath, &app); err != nil {
		return app, err
	}

	return
}

func loadYamlFile(filepath string, data interface{}) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("unable to load configuration file '%s': %s", filepath, err)
	}

	if err = yaml.Unmarshal(content, data); err != nil {
		return fmt.Errorf("unable to parse configuration file '%s': %s", filepath, err)
	}
	return nil
}
