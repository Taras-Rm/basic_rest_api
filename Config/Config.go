package Config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config Cfg

type Cfg struct {
	App App `yaml:"app"`
}

type App struct {
	UserName string `yaml:"username"`
	Pass     string `yaml:"pass"`
	DbName   string `yaml:"dbname"`
	Host     string `yaml:"localhost"`
}

func InitConfig() {
	file, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		panic(err)
	}
}
