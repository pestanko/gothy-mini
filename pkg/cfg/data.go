package cfg

import (
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/user"
	"path"
)

// DataTemplate whole data representation
type DataTemplate struct {
	Users   []user.User     `yaml:"users"`
	Clients []client.Client `yaml:"clients"`
}

// LoadDataTemplate load all the data to the DataTemplate
func LoadDataTemplate(config *AppCfg) (data DataTemplate, err error) {
	filePath := path.Join("config", Vars.Env, config.Data.Load)
	err = loadYamlFile(filePath, &data)
	return
}
