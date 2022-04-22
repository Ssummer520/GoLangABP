package Conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var (
	Configs *Conf
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
}

type Server struct {
	Port         string `yaml:"port"`
	ReadTimeout  string `yaml:"read-timeout"`
	WriteTimeout string `yaml:"write-timeout"`
}

type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"User"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

type Redis struct {
	Addr        string        `yaml:"addr"`
	Pass        string        `yaml:"pass"`
	DB          int           `yaml:"db"`
	Timeout     time.Duration `yaml:"timeout"`
	ExpiredTime int           `yaml:"expired-time"`
}

type RabbitMQ struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	User string `yaml:"User"`
	Pass string `yaml:"pass"`
}

func init() {
	Configs = getConf()
	log.Println("Config init success")
}

func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("demo.Web.Host/Conf/Config.yml")

	if err != nil {
		log.Println("config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("yaml unmarshal error: ", err)
	}
	return c
}
