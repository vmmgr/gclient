package etc

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

type Data struct {
	User  string
	Pass  string
	Token string
	Host  string
}

type BaseData struct {
	Host  string
	User  string
	Pass  string
	Token string
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
	IP string `json:"host"`
}

func GetData(command *cobra.Command) Data {
	base := Base(command)
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	host := config.Host.IP
	token := config.Authentication.Token
	if base.Host != "" {
		host = base.Host
	}
	if base.Token != "" {
		token = base.Token
	}

	return Data{Token: token, Host: host}
}

func Base(cmd *cobra.Command) BaseData {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	token, err := cmd.Flags().GetString("token")
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return BaseData{
		Host:  host,
		Token: token,
	}
}
