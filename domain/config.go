package domain

import (
	"encoding/json"
	"fmt"
)

type Config []Action

type Action struct {
	Action     string          `json:"action"`
	Parameters json.RawMessage `json:"parameters"`
}

func (a *Action) GetParametersBuild() (result ParametersBuild, err error) {
	result = ParametersBuild{}
	err = json.Unmarshal(a.Parameters, &result)
	return
}

func (a *Action) GetParametersPush() (result ParametersPush, err error) {
	result = ParametersPush{}
	err = json.Unmarshal(a.Parameters, &result)
	return
}

func (a *Action) GetParametersCall() (result ParametersCall, err error) {
	result = ParametersCall{}
	err = json.Unmarshal(a.Parameters, &result)
	return
}

type ParametersBuild struct {
	Tag        string `json:"tag"`
	Dockerfile string `json:"dockerfile"`
}

func (p *ParametersBuild) ToArgs() (args []string) {
	args = []string{"docker", "build", "-t", p.Tag}
	if p.Dockerfile != "" {
		args = append(args, fmt.Sprint("--file=", p.Dockerfile))
	}
	args = append(args, ".")
	return
}

type ParametersPush struct {
	Tag string `json:"tag"`
}

type ParametersCall struct {
	Method string `json:"method"`
	URI    string `json:"uri"`
}

func GetExampleConfig() Config {
	c := Config{
		Action{
			Action: "BUILD",
		},
		Action{
			Action: "PUSH",
		},
		Action{
			Action: "CALL",
		},
	}

	bp := ParametersBuild{
		Tag:        "team/repo:version",
		Dockerfile: "",
	}
	c[0].Parameters, _ = json.Marshal(bp)

	pp := ParametersPush{
		Tag: "team/repo:version",
	}
	c[1].Parameters, _ = json.Marshal(pp)
	cp := ParametersCall{
		Method: "POST",
		URI:    "https://buildserver.com/webhooks/x",
	}
	c[2].Parameters, _ = json.Marshal(cp)
	return c
}
