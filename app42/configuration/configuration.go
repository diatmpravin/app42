package configuration

import (
	"encoding/json"
	term "github.com/diatmpravin/app42_client/app42/terminal"
	"io/ioutil"
	"os"
	"os/user"
)

const (
	filePermissions = 0644
	dirPermissions  = 0700
)

type Keys struct {
	ApiKey     string
	SecretKey  string
	ApiVersion string
}

func DeleteKeys() (err error) {
	file, err := configFile()

	if err != nil {
		return
	}

	os.Remove(file)
	return
}

func (k Keys) Save() (err error) {
	bytes, err := json.Marshal(k)

	if err != nil {
		return
	}

	file, err := configFile()

	if err != nil {
		return
	}

	err = ioutil.WriteFile(file, bytes, filePermissions)

	return
}

func configFile() (file string, err error) {
	currentUser, err := user.Current()

	if err != nil {
		return
	}

	configDir := currentUser.HomeDir + "/.app42"

	err = os.MkdirAll(configDir, dirPermissions)

	if err != nil {
		return
	}

	file = configDir + "/app42paas.yml"

	return
}

func Load() (k Keys, err error) {
	file, err := configFile()

	if err != nil {
		return
	}

	data, err := ioutil.ReadFile(file)

	if err != nil {
		term.Failed("API key and Secret key not found", err)
		return
	}

	err = json.Unmarshal(data, &k)

	return
}

func ShowKeys() {
	config, err := Load()

	if err != nil {
		term.Failed("File is invalid", err)
		return

	}

	showConfiguration(config)
}

func showConfiguration(k Keys) {
	term.Say(term.Yellow("=== API/Secret Key ==="))
	term.Say("API Key:    %s ", k.ApiKey)
	term.Say("Secret Key: %s ", k.SecretKey)
}
