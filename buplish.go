package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/just1689/buplish/domain"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var configFile = flag.String("config", "buplish.json", "Set the file for buplish to use")
var initFile = flag.Bool("init", false, "Set to true to generate an empty config")

func main() {
	flag.Parse()

	if *initFile {
		generateFile()
		return
	}

	config := loadConfig()
	for a, b := range config {
		fmt.Sprintln("For rule %i action is %s", a, b.Action)
	}

}

func generateFile() {
	c := domain.Config{
		domain.Action{
			Action: "BUILD",
		},
		domain.Action{
			Action: "PUSH",
		},
		domain.Action{
			Action: "CALL",
		},
	}

	bp := domain.ParametersBuild{
		Tag:        "team/repo:version",
		Dockerfile: "",
	}
	c[0].Parameters, _ = json.Marshal(bp)

	pp := domain.ParametersPush{
		Tag: "team/repo:version",
	}
	c[1].Parameters, _ = json.Marshal(pp)
	cp := domain.ParametersCall{
		Method: "POST",
		URI:    "https://buildserver.com/webhooks/x",
	}
	c[2].Parameters, _ = json.Marshal(cp)

	b, err := json.Marshal(c)
	if err != nil {
		logrus.Panic(err)
	}
	ioutil.WriteFile(*configFile, b, 0600)

}

func loadConfig() domain.Config {
	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		logrus.Panic(err)
	}
	result := domain.Config{}
	if err = json.Unmarshal(b, &result); err != nil {
		logrus.Panic("Could not unmarshal with err ", err.Error())
	}
	return result

}
