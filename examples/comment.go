package comment

import (
	"code.google.com/p/gopass"
	"flag"
	"fmt"
	"github.com/jbhat/go-jira-client"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
	"time"
)

type Config struct {
	Host         string `yaml:"host"`
	ApiPath      string `yaml:"api_path"`
	ActivityPath string `yaml:"activity_path"`
	Login        string `yaml:"login"`
}

func TestComment() {
	startedAt := time.Now()
	defer func() {
		fmt.Printf("processed in %v\n", time.Now().Sub(startedAt))
	}()

	help := flag.Bool("help", false, "Show usage")

	// read config file
	file, e := ioutil.ReadFile("config.yml")
	if e != nil {
		fmt.Printf("Config file error: %v\n", e)
		os.Exit(1)
	}

	// parse config file
	config := new(Config)
	err := goyaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	pass, err := gopass.GetPass("Password: ")
	if err != nil {
		panic(err.Error())
	}

	jira := gojira.NewJira(
		config.Host,
		config.ApiPath,
		config.ActivityPath,
		&gojira.Auth{config.Login, pass},
	)

	var i string
	flag.StringVar(&i, "i", "", "Issue key")

	var comment string
	flag.StringVar(&comment, "c", "", "Comment to be added")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || i == "" || comment == "" {
		flag.Usage()
		return
	}

	issue, err := jira.Issue(i)
	if err != nil {
		panic(err.Error())
	}
	err = jira.AddComment(&issue, comment)
	if err != nil {
		panic(err.Error())
	}
}
