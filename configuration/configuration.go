package configuration

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
var configFilePath string

/*
	Service configuration struct
*/
var conf Configuration

func Configs() Configuration {
	return conf
}

type Configuration struct {
	// HAProxy output configuration
	HAProxy HAProxy
}

type HAProxy struct {
	TemplatePath            string
	OutputPath              string
	ReloadCommand           string
	ReloadValidationCommand string
	ReloadCleanupCommand    string
}

func init() {
	log.Println("initialized config")
	flag.StringVar(&configFilePath, "config", "config/development.json", "Full path of the configuration JSON file")
	err := FromFile(configFilePath, &conf)
	if err != nil {
		log.Fatal(err)
	}
}

func (config *Configuration) FromFile(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return json.Unmarshal(content, &config)
}

func FromFile(filePath string, conf *Configuration) error {
	err := conf.FromFile(filePath)
	setValueFromEnv(&conf.HAProxy.TemplatePath, "HAPROXY_TEMPLATE_PATH")
	setValueFromEnv(&conf.HAProxy.OutputPath, "HAPROXY_OUTPUT_PATH")
	setValueFromEnv(&conf.HAProxy.ReloadCommand, "HAPROXY_RELOAD_CMD")
	setValueFromEnv(&conf.HAProxy.ReloadValidationCommand, "HAPROXY_RELOAD_VALIDATION_CMD")
	setValueFromEnv(&conf.HAProxy.ReloadCleanupCommand, "HAPROXY_RELOAD_CLEANUP_CMD")

	return err
}

func setValueFromEnv(field *string, envVar string) {
	env := os.Getenv(envVar)
	if len(env) > 0 {
		log.Printf("Using environment override %s=%s", envVar, env)
		*field = env
	}
}
