package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"flag"
	"github.com/golang/glog"
	"regexp"
)

type Config struct {
	PathPrefix string			`yaml:"pathPrefix"`
	Changes map[string]Changes		`yaml:"changes"`
}

type Changes struct {
	Variables map[string][]string	`yaml:"variables"`
	Types map[string][]string	`yaml:"types"`
}

var (
	cfg = &Config{}
	configFlag string
)

func init() {

	flag.StringVar(&configFlag, "config", "", "configFile")

	flag.Parse()

	if configFlag != "" {
		yamlFile, err := ioutil.ReadFile(configFlag)
		if err != nil {
			glog.Fatalf("Cannot get config file %s Get err   #%v ", configFlag, err)
			os.Exit(-1)
		}
		if err != nil {
			glog.Fatalf("Config parse error: %v", err)
			os.Exit(-1)
		}
		err = yaml.Unmarshal(yamlFile,&cfg)
		if err != nil {
			glog.Fatalf("Config parse error: %v", err)
			os.Exit(-1)
		}
	}else{
		glog.Fatalf("Need a config file")
		os.Exit(-1)
	}

	if cfg == nil {
		glog.Fatalf("Config file can not be empty")
		os.Exit(-1)
	}
}

func main() {
	for path, changes := range cfg.Changes {
		fullPath := filepath.Join(cfg.PathPrefix,path)
		read, err := ioutil.ReadFile(fullPath)
		if err != nil {
			panic(err)
		}
		fmt.Println(fullPath)
		newContent := string(read)
		if changes.Variables != nil {
			for variable, annotations := range changes.Variables{
				for _, annotation := range annotations {
					m := regexp.MustCompile("(\\t+)(" + variable + ".*)") 
					newContent = m.ReplaceAllString(newContent, "${1}" + annotation + "\n${1}${2}")
				}
			}

		}
		if changes.Types != nil {
			for types, annotations := range changes.Types{
				for _, annotation := range annotations {
					m := regexp.MustCompile("(\\t*)(type " + types + " struct.*)") 
					newContent = m.ReplaceAllString(newContent, "${1}" + annotation + "\n${1}${2}")
				}
			}

		}
		err = ioutil.WriteFile(fullPath, []byte(newContent), 0)
		if err != nil {
			panic(err)
		}

	}
}
