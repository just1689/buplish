package disk

import (
	"encoding/json"
	"github.com/just1689/buplish/domain"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func GenerateFile(configFile *string) {
	c := domain.GetExampleConfig()
	b, err := json.Marshal(c)
	if err != nil {
		logrus.Panic(err)
	}
	ioutil.WriteFile(*configFile, b, 0600)

}

func LoadConfig(configFile *string) domain.Config {
	logrus.Println("Loading", *configFile)
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
