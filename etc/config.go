package etc

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/vmmgr/gclient/cmd"
	"io/ioutil"
)

type Data struct {
	User  string
	Pass  string
	Token string
	Host  string
}

type Config struct {
	Authentication AuthenticationData `json:authentication`
	Host           HostData           `json:host`
}

type AuthenticationData struct {
	User  string `json:"user"`
	Token string `json:"pass"`
}
type HostData struct {
	Host string `json:"host"`
}

func GetData(command *cobra.Command) Data {
	base := cmd.Base(command)
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	host := config.Host.Host
	user := config.Authentication.User
	pass := base.Pass
	token := config.Authentication.Token
	if base.Host != "" {
		host = base.Host
	}
	if base.User != "" {
		host = base.User
	}
	if base.Token != "" {
		host = base.Token
	}

	return Data{User: user, Pass: pass, Token: token, Host: host}
}
