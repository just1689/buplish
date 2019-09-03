package main

import (
	"flag"
	"fmt"
	"github.com/just1689/buplish/disk"
	"github.com/just1689/buplish/domain"
	"github.com/sirupsen/logrus"
	"io"
	"os/exec"
)

var configFile = flag.String("config", "buplish.json", "Set the file for buplish to use")
var initFile = flag.Bool("init", false, "Set to true to generate an empty config")

var actionHandlers = map[string]func(action domain.Action, num int){
	"BUILD": handleBuild,
	"PUSH": func(action domain.Action, num int) {
		fmt.Println("push...")
	},
	"CALL": func(action domain.Action, num int) {
		fmt.Println("call...")
	},
}

func main() {
	o, _ := exec.Command("docker", "build", "-t", "team/repo:version", ".").StdoutPipe()
	p := make([]byte, 256)
	for {
		n, err := o.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}

	if true {
		return
	}

	flag.Parse()
	if *initFile {
		disk.GenerateFile(configFile)
		return
	}

	config := disk.LoadConfig(configFile)
	logrus.Println("Configuration loaded. Running tasks", len(config))
	for num, action := range config {
		f, found := actionHandlers[action.Action]
		if !found {
			logrus.Fatalln("unrecognized action:", action.Action)
		}
		f(action, num)
	}

}

func handleBuild(action domain.Action, num int) {
	detail, err := action.GetParametersBuild()
	if err != nil {
		logrus.Errorln("Failed at ", num, action.Action)
		logrus.Panic(err)
	}
	args := detail.ToArgs()
	fmt.Println(args)
	c := exec.Command("cmd.exe", args...)
	if err != nil {
		logrus.Panic(err)
	}

	if err := c.Run(); err != nil {
		logrus.Errorln(err)
	}

}
