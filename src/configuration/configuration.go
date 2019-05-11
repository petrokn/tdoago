package configuration

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	CoreConfiguration   CoreConfiguration   `yaml:"core"`
	ServerConfiguration ServerConfiguration `yaml:"server"`
	ClientConfiguration CoreConfiguration   `yaml:"client"`
	Messaging           Messaging           `yaml:"messaging"`
}

type CoreConfiguration struct {
	Modeling           bool    `yaml:"modeling"`
	LogFile            string  `yaml:"log_file"`
	LocalizationRadius float64 `yaml:"localization_radius"`
	ThreadAmount       int8    `yaml:"thread_amount"`
}

type ServerConfiguration struct {
	ServerAddress      string `yaml:"server_address"`
	ServerPort         int    `yaml:"server_port"`
	LocalizationTrials int8   `yaml:"localization_trials"`
}

type ClientConfiguration struct {
	AudioPath     string `yaml:"audio_path"`
	ProxiesAmount int    `yaml:"proxies_amount"`
}

type Messaging struct {
	Hosts string `yaml:"hosts"`
}

func Configure(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ConfigFileReadError %v ", err)
	}

	var config *Config

	err = yaml.Unmarshal(yamlFile, &config)
	log.Println(config)
	if err != nil {
		log.Fatalf("ConfigFileUnmarshalError %v", err)
	}

	return config
}
