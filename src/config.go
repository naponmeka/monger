package src

import (
	"encoding/json"
	"path/filepath"

	"github.com/naponmeka/robone/list"
	"github.com/shibukawa/configdir"
)

type Config struct {
	Timeout          int             `json:"timeout"`
	DefaultView      int             `json:"default-view"`
	SavedConnections []list.ListItem `json:"saved-connections"`
}

func readConfig() (config Config) {
	configDirs := configdir.New("com.robone.www", "robone")
	configDirs.LocalPath, _ = filepath.Abs(".")
	folder := configDirs.QueryFolderContainsFile("setting.json")
	if folder != nil {
		data, _ := folder.ReadFile("setting.json")
		json.Unmarshal(data, &config)
	} else {
		config.Timeout = 15
	}
	return
}

func saveConfig(config Config) {
	configDirs := configdir.New("com.robone.www", "robone")
	data, _ := json.Marshal(&config)
	folders := configDirs.QueryFolders(configdir.Global)
	folders[0].WriteFile("setting.json", data)
}
