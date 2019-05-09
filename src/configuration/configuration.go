package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	CoreConfiguration CoreConfiguration
	ServerConfiguration ServerConfiguration
	ClientConfiguration CoreConfiguration
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

func Configure(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var config *Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
