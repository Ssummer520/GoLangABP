package Conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	User        string `yaml:"user"`
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
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

func init() {
	Configs = getConf()
	klog.Info(Configs, "Config init success")
}

func getConf() *Conf {
	var c *Conf
	klog.Info(GetAppPath())
	file, err := ioutil.ReadFile("./demo.Web.Host/conf/Config.yml")
	if err != nil {
		klog.Error(err)
	}
	klog.Error(c)
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		klog.Error(err)
	}
	return c
}
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
