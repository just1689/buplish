package domain

import "encoding/json"

type Config []Action

type Action struct {
	Action     string          `json:"action"`
	Parameters json.RawMessage `json:"parameters"`
}

type ParametersBuild struct {
	Tag        string `json:"tag"`
	Dockerfile string `json:"dockerfile"`
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
